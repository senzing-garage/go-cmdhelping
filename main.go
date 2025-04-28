/*
 */
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/senzing-garage/go-cmdhelping/cmdhelper"
	"github.com/senzing-garage/go-cmdhelping/constant"
	"github.com/senzing-garage/go-cmdhelping/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	Use   string = "Name of command"
	Short string = "Short description of command"
	Long  string = "Long description of command."
)

func Run(cmd *cobra.Command, args []string) {
	_ = cmd
	_ = args

	outputf(
		"--%-12s  %s: %s\n",
		option.DatabaseURL.Arg,
		fmt.Sprintf(option.DatabaseURL.Help, option.DatabaseURL.Envar),
		viper.GetString(option.DatabaseURL.Arg),
	)
	outputf(
		"--%-12s  %s: %d\n",
		option.HTTPPort.Arg,
		fmt.Sprintf(option.HTTPPort.Help, option.HTTPPort.Envar),
		viper.GetInt(option.HTTPPort.Arg),
	)
	outputf(
		"--%-12s  %s: %s\n",
		option.LogLevel.Arg,
		fmt.Sprintf(option.LogLevel.Help, option.LogLevel.Envar),
		viper.GetString(option.LogLevel.Arg),
	)
}

func main() {
	// Define the command-line options / environment variables used for context variables.
	ContextVariables := []option.ContextVariable{option.DatabaseURL, option.HTTPPort, option.LogLevel}

	// Initialize viper.

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.SetEnvPrefix(constant.SetEnvPrefix)

	// Initialize cobra.

	cobraCommand := &cobra.Command{
		Use:     Use,
		Short:   Short,
		Long:    Long,
		Run:     Run,
		Version: cmdhelper.Version("1234", "5"),
	}

	// Used in local initialization.

	cmdhelper.Init(cobraCommand, ContextVariables)

	// Usually used in cobra.Command.PreRun.

	cmdhelper.PreRun(cobraCommand, os.Args, Use, ContextVariables)

	// Execute the cobra command which prints the output from "Run:" function.

	err := cobraCommand.Execute()
	if err != nil {
		panic(err)
	}
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

func outputf(format string, message ...any) {
	fmt.Printf(format, message...) //nolint
}
