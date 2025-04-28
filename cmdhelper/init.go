package cmdhelper

import (
	"fmt"

	"github.com/senzing-garage/go-cmdhelping/option"
	"github.com/senzing-garage/go-cmdhelping/option/optiontype"
	"github.com/spf13/cobra"
)

// Performs cobra.Flags.<datatype> operations on a list of contextVariables.
func Init(cobraCommand *cobra.Command, contextVariables []option.ContextVariable) {
	for _, contextVariable := range contextVariables {
		SetCobraFlag(cobraCommand, contextVariable)
	}
}

// Performs cobra.Flag.<datatype> on a option.ContextVariable.
func SetCobraFlag(cobraCommand *cobra.Command, contextVariable option.ContextVariable) {
	helpMessage := fmt.Sprintf(contextVariable.Help, contextVariable.Envar)

	switch contextVariable.Type {
	case optiontype.Bool:
		value, isOK := contextVariable.Default.(bool)
		if isOK {
			cobraCommand.Flags().Bool(contextVariable.Arg, value, helpMessage)
		} else {
			panic(contextVariable.Arg + " should be boolean")
		}
	case optiontype.Int:
		value, isOK := contextVariable.Default.(int)
		if isOK {
			cobraCommand.Flags().Int(contextVariable.Arg, value, helpMessage)
		} else {
			panic(contextVariable.Arg + " should be int")
		}
	case optiontype.String:
		value, isOK := contextVariable.Default.(string)
		if isOK {
			cobraCommand.Flags().String(contextVariable.Arg, value, helpMessage)
		} else {
			panic(contextVariable.Arg + " should be string")
		}
	case optiontype.StringSlice:
		value, isOK := contextVariable.Default.([]string)
		if isOK {
			cobraCommand.Flags().StringSlice(contextVariable.Arg, value, helpMessage)
		} else {
			panic(contextVariable.Arg + " should be string")
		}
	}
}
