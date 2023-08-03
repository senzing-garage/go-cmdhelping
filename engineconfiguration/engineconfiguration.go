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
		options := map[string]string{
			"configPath":          aViper.GetString(option.ConfigPath.Arg),
			"databaseUrl":         aViper.GetString(option.DatabaseUrl.Arg),
			"licenseStringBase64": aViper.GetString(option.LicenseStringBase64.Arg),
			"resourcePath":        aViper.GetString(option.ResourcePath.Arg),
			"senzingDirectory":    aViper.GetString(option.SenzingDirectory.Arg),
			"supportPath":         aViper.GetString(option.SupportPath.Arg),
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
