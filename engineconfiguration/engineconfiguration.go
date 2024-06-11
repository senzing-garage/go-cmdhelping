package engineconfiguration

import (
	"context"

	"github.com/senzing-garage/go-cmdhelping/option"
	"github.com/senzing-garage/go-helpers/engineconfigurationjson"
	"github.com/spf13/viper"
)

// Given variables in Viper, construct the Senzing engine configuration JSON.
func BuildSenzingEngineConfigurationJSON(ctx context.Context, aViper *viper.Viper) (string, error) {
	_ = ctx
	var err error
	result := aViper.GetString(option.EngineConfigurationJSON.Arg)
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
		result, err = engineconfigurationjson.BuildSimpleSystemConfigurationJsonUsingMap(options)
	}
	return result, err
}

// Convenience method for engineconfigurationjson.VerifySenzingEngineConfigurationJSON
func VerifySenzingEngineConfigurationJSON(ctx context.Context, engineConfigurationJSON string) error {
	return engineconfigurationjson.VerifySenzingEngineConfigurationJson(ctx, engineConfigurationJSON)
}

// Given variables in Viper, construct and verify the Senzing engine configuration JSON.
func BuildAndVerifySenzingEngineConfigurationJSON(ctx context.Context, aViper *viper.Viper) (string, error) {
	senzingEngineConfigurationJSON, err := BuildSenzingEngineConfigurationJSON(ctx, aViper)
	if err != nil {
		return senzingEngineConfigurationJSON, err
	}
	err = VerifySenzingEngineConfigurationJSON(ctx, senzingEngineConfigurationJSON)
	return senzingEngineConfigurationJSON, err
}
