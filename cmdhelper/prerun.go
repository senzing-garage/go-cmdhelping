package cmdhelper

import (
	"fmt"
	"os"
	"strings"

	"github.com/senzing-garage/go-cmdhelping/constant"
	"github.com/senzing-garage/go-cmdhelping/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// If a configuration file is present, load it.
func loadConfigurationFile(cobraCommand *cobra.Command, configName string) {
	var configuration string

	configFlag := cobraCommand.Flags().Lookup(option.Configuration.Arg)
	if configFlag != nil {
		configuration = configFlag.Value.String()
	}

	if configuration != "" { // Use configuration file specified as a command line option.
		viper.SetConfigFile(configuration)
	} else { // Search for a configuration file.
		// Determine home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Specify configuration file name.

		viper.SetConfigName(configName)
		viper.SetConfigType("yaml")

		// Define search path order.

		viper.AddConfigPath(home + "/.senzing-tools")
		viper.AddConfigPath(home)
		viper.AddConfigPath("/etc/senzing-tools")
	}

	// If a config file is found, read it in.

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Applying configuration file:", viper.ConfigFileUsed())
	}
}

// Configure Viper with user-specified options.
func loadOptions(cobraCommand *cobra.Command, contextVariables []option.ContextVariable) {
	var err error

	viper.AutomaticEnv()

	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix(constant.SetEnvPrefix)

	for _, contextVar := range contextVariables {
		viper.SetDefault(contextVar.Arg, contextVar.Default)

		err = viper.BindPFlag(contextVar.Arg, cobraCommand.Flags().Lookup(contextVar.Arg))
		if err != nil {
			panic(err)
		}
	}
}

// A function that can be wrapped for use by cobra.Command.PreRun.
func PreRun(cobraCommand *cobra.Command, args []string, configName string, contextVariables []option.ContextVariable) {
	_ = args

	loadConfigurationFile(cobraCommand, configName)
	loadOptions(cobraCommand, contextVariables)
	cobraCommand.SetVersionTemplate(constant.VersionTemplate)
}
