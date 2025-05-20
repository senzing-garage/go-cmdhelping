package settings

import (
	"context"

	"github.com/senzing-garage/go-cmdhelping/option"
	"github.com/senzing-garage/go-helpers/settings"
	"github.com/senzing-garage/go-helpers/wraperror"
	"github.com/spf13/viper"
)

// Given variables in Viper, construct the Senzing engine configuration JSON.
func BuildSettings(ctx context.Context, aViper *viper.Viper) (string, error) {
	_ = ctx

	var err error

	result := aViper.GetString(option.EngineSettings.Arg)
	if len(result) == 0 {
		optionsList := map[string]option.ContextVariable{
			"configPath":          option.ConfigPath,
			"databaseUrl":         option.DatabaseURL,
			"licenseStringBase64": option.LicenseStringBase64,
			"resourcePath":        option.ResourcePath,
			"senzingDirectory":    option.SenzingDirectory,
			"supportPath":         option.SupportPath,
		}
		options := map[string]string{}

		for key, contextVariable := range optionsList {
			if aViper.IsSet(contextVariable.Arg) {
				if aViper.GetString(contextVariable.Arg) != contextVariable.Default {
					options[key] = aViper.GetString(contextVariable.Arg)
				}
			}
		}

		result, err = settings.BuildSimpleSettingsUsingMap(options)
	}

	return result, wraperror.Errorf(err, wraperror.NoMessage)
}

// Convenience method for engineconfigurationjson.VerifySettings.
func VerifySettings(ctx context.Context, theSettings string) error {
	err := settings.VerifySettings(ctx, theSettings)

	return wraperror.Errorf(err, "go-helper.settings.VerifySettings")
}

// Given variables in Viper, construct and verify the Senzing engine configuration JSON.
func BuildAndVerifySettings(ctx context.Context, aViper *viper.Viper) (string, error) {
	theSettings, err := BuildSettings(ctx, aViper)
	if err != nil {
		return theSettings, wraperror.Errorf(err, "BuildSettings")
	}

	err = VerifySettings(ctx, theSettings)

	return theSettings, wraperror.Errorf(err, wraperror.NoMessage)
}
