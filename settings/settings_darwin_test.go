//go:build darwin

package settings

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/senzing-garage/go-cmdhelping/constant"
	"github.com/senzing-garage/go-cmdhelping/option"
	"github.com/senzing-garage/go-helpers/settings"
	"github.com/spf13/viper"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestBuildAndVerifySettings(test *testing.T) {
	_ = test
	ctx := context.TODO()
	senzingPath := settings.GetSenzingPath()
	var contextVariables = []option.ContextVariable{
		option.ConfigPath.SetDefault(fmt.Sprintf("%s/er/etc", senzingPath)),
		option.DatabaseURL.SetDefault("sqlite3://na:na@/tmp/sqlite/G2C.db"),
		option.LicenseStringBase64.SetDefault("ABCD12134"),
		option.ResourcePath.SetDefault(fmt.Sprintf("%s/er/resources", senzingPath)),
		option.SenzingDirectory,
		option.SupportPath.SetDefault(fmt.Sprintf("%s/data", senzingPath)),
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
