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
	switch contextVariable.Type {
	case optiontype.Bool:
		setCobraFlagBool(cobraCommand, contextVariable)
	case optiontype.Int:
		setCobraFlagInt(cobraCommand, contextVariable)
	case optiontype.String:
		setCobraFlagString(cobraCommand, contextVariable)
	case optiontype.StringSlice:
		setCobraFlagStringSlice(cobraCommand, contextVariable)
	case optiontype.Uint:
		setCobraFlagUint(cobraCommand, contextVariable)
	case optiontype.Uint32:
		setCobraFlagUint32(cobraCommand, contextVariable)
	case optiontype.Uint64:
		setCobraFlagUint64(cobraCommand, contextVariable)
	}
}

func setCobraFlagBool(cobraCommand *cobra.Command, contextVariable option.ContextVariable) {
	helpMessage := fmt.Sprintf(contextVariable.Help, contextVariable.Envar)
	value, isOK := contextVariable.Default.(bool)

	if isOK {
		cobraCommand.Flags().Bool(contextVariable.Arg, value, helpMessage)
	} else {
		panic(contextVariable.Arg + " should be boolean")
	}
}

func setCobraFlagInt(cobraCommand *cobra.Command, contextVariable option.ContextVariable) {
	helpMessage := fmt.Sprintf(contextVariable.Help, contextVariable.Envar)
	value, isOK := contextVariable.Default.(int)

	if isOK {
		cobraCommand.Flags().Int(contextVariable.Arg, value, helpMessage)
	} else {
		panic(contextVariable.Arg + " should be int")
	}
}

func setCobraFlagString(cobraCommand *cobra.Command, contextVariable option.ContextVariable) {
	helpMessage := fmt.Sprintf(contextVariable.Help, contextVariable.Envar)
	value, isOK := contextVariable.Default.(string)

	if isOK {
		cobraCommand.Flags().String(contextVariable.Arg, value, helpMessage)
	} else {
		panic(contextVariable.Arg + " should be string")
	}
}

func setCobraFlagStringSlice(cobraCommand *cobra.Command, contextVariable option.ContextVariable) {
	helpMessage := fmt.Sprintf(contextVariable.Help, contextVariable.Envar)
	value, isOK := contextVariable.Default.([]string)

	if isOK {
		cobraCommand.Flags().StringSlice(contextVariable.Arg, value, helpMessage)
	} else {
		panic(contextVariable.Arg + " should be string")
	}
}

func setCobraFlagUint(cobraCommand *cobra.Command, contextVariable option.ContextVariable) {
	helpMessage := fmt.Sprintf(contextVariable.Help, contextVariable.Envar)
	value, isOK := contextVariable.Default.(uint)

	if isOK {
		cobraCommand.Flags().Uint(contextVariable.Arg, value, helpMessage)
	} else {
		panic(contextVariable.Arg + " should be uint")
	}
}

func setCobraFlagUint32(cobraCommand *cobra.Command, contextVariable option.ContextVariable) {
	helpMessage := fmt.Sprintf(contextVariable.Help, contextVariable.Envar)
	value, isOK := contextVariable.Default.(uint32)

	if isOK {
		cobraCommand.Flags().Uint32(contextVariable.Arg, value, helpMessage)
	} else {
		panic(contextVariable.Arg + " should be uint32")
	}
}

func setCobraFlagUint64(cobraCommand *cobra.Command, contextVariable option.ContextVariable) {
	helpMessage := fmt.Sprintf(contextVariable.Help, contextVariable.Envar)
	value, isOK := contextVariable.Default.(uint64)

	if isOK {
		cobraCommand.Flags().Uint64(contextVariable.Arg, value, helpMessage)
	} else {
		panic(contextVariable.Arg + " should be uint64")
	}
}
