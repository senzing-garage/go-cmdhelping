//go:build linux

package settings_test

import (
	"strings"
	"testing"

	"github.com/senzing-garage/go-cmdhelping/constant"
	"github.com/senzing-garage/go-cmdhelping/option"
	"github.com/senzing-garage/go-cmdhelping/settings"
	helperSettings "github.com/senzing-garage/go-helpers/settings"
	"github.com/spf13/viper"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestBuildAndVerifySettings(test *testing.T) { //nolint
	ctx := test.Context()
	senzingPath := helperSettings.GetSenzingPath()

	contextVariables := []option.ContextVariable{
		option.ConfigPath.SetDefault("/etc/opt/senzing"),
		option.DatabaseURL.SetDefault("sqlite3://na:na@/tmp/sqlite/G2C.db"),
		option.LicenseStringBase64.SetDefault("ABCD12134"),
		option.ResourcePath.SetDefault(senzingPath + "/er/resources"),
		option.SenzingDirectory,
		option.SupportPath.SetDefault(senzingPath + "/data"),
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.SetEnvPrefix(constant.SetEnvPrefix)

	for _, contextVariable := range contextVariables {
		viper.SetDefault(contextVariable.Arg, contextVariable.Default)
	}

	_, err := settings.BuildAndVerifySettings(ctx, viper.GetViper())
	if err != nil {
		panic(err)
	}
}
