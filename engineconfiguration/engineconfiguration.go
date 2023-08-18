package engineconfiguration

import (
	"context"

	"github.com/senzing/go-cmdhelping/option"
	"github.com/senzing/go-common/g2engineconfigurationjson"
	"github.com/spf13/viper"
)

// Given variables in Viper, construct the Senzing engine configuration JSON.
func BuildSenzingEngineConfigurationJson(ctx context.Context, aViper *viper.Viper) (string, error) {
	var err error = nil
	result := aViper.GetString(option.EngineConfigurationJson.Arg)
	if len(result) == 0 {
		optionsList := map[string]option.ContextVariable{
			"configPath":          option.ConfigPath,
			"databaseUrl":         option.DatabaseUrl,
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
		result, err = g2engineconfigurationjson.BuildSimpleSystemConfigurationJsonUsingMap(options)
	}
	return result, err
}

// Convenience method for g2engineconfigurationjson.VerifySenzingEngineConfigurationJson
func VerifySenzingEngineConfigurationJson(ctx context.Context, engineConfigurationJson string) error {
	return g2engineconfigurationjson.VerifySenzingEngineConfigurationJson(ctx, engineConfigurationJson)
}

// Given variables in Viper, construct and verify the Senzing engine configuration JSON.
func BuildAndVerifySenzingEngineConfigurationJson(ctx context.Context, aViper *viper.Viper) (string, error) {
	senzingEngineConfigurationJson, err := BuildSenzingEngineConfigurationJson(ctx, aViper)
	if err != nil {
		return senzingEngineConfigurationJson, err
	}
	err = VerifySenzingEngineConfigurationJson(ctx, senzingEngineConfigurationJson)
	return senzingEngineConfigurationJson, err
}
