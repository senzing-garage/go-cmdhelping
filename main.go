/*
 */
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/senzing/go-cmdhelping/cmdhelper"
	"github.com/senzing/go-cmdhelping/constant"
	"github.com/senzing/go-cmdhelping/option"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	Use   string = "Name of command"
	Short string = "Short description of command"
	Long  string = "Long description of command."
)

func main() {

	// Define the command-line options / environment variables used for context variables.

	var ContextVariables = []option.ContextVariable{
		option.DatabaseUrl,
		option.HttpPort,
		option.LogLevel,
	}

	// Initialize viper.

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.SetEnvPrefix(constant.SetEnvPrefix)

	// Initialize cobra.

	var cobraCommand = &cobra.Command{
		Use:   Use,
		Short: Short,
		Long:  Long,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("--%-12s  %s: %s\n", option.DatabaseUrl.Arg, fmt.Sprintf(option.DatabaseUrl.Help, option.DatabaseUrl.Envar), viper.GetString(option.DatabaseUrl.Arg))
			fmt.Printf("--%-12s  %s: %d\n", option.HttpPort.Arg, fmt.Sprintf(option.HttpPort.Help, option.HttpPort.Envar), viper.GetInt(option.HttpPort.Arg))
			fmt.Printf("--%-12s  %s: %s\n", option.LogLevel.Arg, fmt.Sprintf(option.LogLevel.Help, option.LogLevel.Envar), viper.GetString(option.LogLevel.Arg))
		},
		Version: cmdhelper.Version("1234", "5"),
	}

	// Used in local initialization.

	cmdhelper.Init(cobraCommand, ContextVariables)

	// Usually used in cobra.Command.PreRun.

	cmdhelper.PreRun(cobraCommand, os.Args, Use, ContextVariables)

	// Execute the cobra command which prints the output from "Run:" function.

	cobraCommand.Execute()
}
