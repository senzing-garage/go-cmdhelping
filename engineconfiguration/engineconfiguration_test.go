package engineconfiguration

import (
	"context"
	"strings"
	"testing"

	"github.com/senzing-garage/go-cmdhelping/constant"
	"github.com/senzing-garage/go-cmdhelping/option"
	"github.com/spf13/viper"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestBuildSenzingEngineConfigurationJson(test *testing.T) {
	ctx := context.TODO()

	var contextVariables = []option.ContextVariable{
		option.ConfigPath.SetDefault("/tmp/ConfigPath"),
		option.DatabaseUrl.SetDefault("sqlite3://na:na@/tmp/sqlite/G2C.db"),
		option.LicenseStringBase64.SetDefault("ABCD12134"),
		option.ResourcePath.SetDefault("/tmp/ResourcePath"),
		option.SenzingDirectory,
		option.SupportPath.SetDefault("/tmp/SupportPath"),
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.SetEnvPrefix(constant.SetEnvPrefix)

	for _, contextVariable := range contextVariables {
		viper.SetDefault(contextVariable.Arg, contextVariable.Default)
	}

	_, err := BuildSenzingEngineConfigurationJson(ctx, viper.GetViper())
	if err != nil {
		panic(err)
	}
}
