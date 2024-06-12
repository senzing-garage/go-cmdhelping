//go:build darwin

package settings

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

func TestBuildAndVerifySettings(test *testing.T) {
	_ = test
	ctx := context.TODO()

	var contextVariables = []option.ContextVariable{
		option.ConfigPath.SetDefault("/opt/senzing/g2/etc"),
		option.DatabaseURL.SetDefault("sqlite3://na:na@/tmp/sqlite/G2C.db"),
		option.LicenseStringBase64.SetDefault("ABCD12134"),
		option.ResourcePath.SetDefault("/opt/senzing/g2/resources"),
		option.SenzingDirectory,
		option.SupportPath.SetDefault("/opt/senzing/g2/data"),
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.SetEnvPrefix(constant.SetEnvPrefix)

	for _, contextVariable := range contextVariables {
		viper.SetDefault(contextVariable.Arg, contextVariable.Default)
	}

	_, err := BuildAndVerifySettings(ctx, viper.GetViper())
	if err != nil {
		panic(err)
	}
}
