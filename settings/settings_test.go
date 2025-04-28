package settings_test

import (
	"strings"
	"testing"

	"github.com/senzing-garage/go-cmdhelping/constant"
	"github.com/senzing-garage/go-cmdhelping/option"
	"github.com/senzing-garage/go-cmdhelping/settings"
	"github.com/spf13/viper"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestBuildSettings(test *testing.T) {
	test.Parallel()
	ctx := test.Context()

	contextVariables := []option.ContextVariable{
		option.ConfigPath.SetDefault("/tmp/ConfigPath"),
		option.DatabaseURL.SetDefault("sqlite3://na:na@/tmp/sqlite/G2C.db"),
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

	_, err := settings.BuildSettings(ctx, viper.GetViper())
	if err != nil {
		panic(err)
	}
}
