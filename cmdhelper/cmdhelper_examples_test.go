package cmdhelper_test

import (
	"fmt"

	"github.com/senzing-garage/go-cmdhelping/cmdhelper"
	"github.com/senzing-garage/go-cmdhelping/option"
	"github.com/senzing-garage/go-cmdhelping/option/optiontype"
	"github.com/spf13/cobra"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleInit() {
	cobraCommand := &cobra.Command{
		Use:   "example-use",
		Short: "example-short",
		Long:  `example-long`,
	}

	contextVariables := []option.ContextVariable{
		{
			Default: "",
			Envar:   "MY_VARIABLE",
			Help:    "Description of my variable [%s]",
			Arg:     "my-variable",
			Type:    optiontype.String,
		},
	}

	cmdhelper.Init(cobraCommand, contextVariables)
	// Output:
}

func ExampleOsLookupEnvBool() {
	fmt.Println(option.OsLookupEnvBool("NOT_AN_ENVIRONMENT_VARIABLE", true))
	// Output: true
}

func ExampleOsLookupEnvInt() {
	fmt.Println(option.OsLookupEnvInt("NOT_AN_ENVIRONMENT_VARIABLE", 10))
	// Output: 10
}

func ExampleOsLookupEnvString() {
	fmt.Println(option.OsLookupEnvString("NOT_AN_ENVIRONMENT_VARIABLE", "default"))
	// Output: default
}

func ExamplePreRun() {
	cobraCommand := &cobra.Command{
		Use:   "example-use",
		Short: "example-short",
		Long:  `example-long`,
	}

	contextVariables := []option.ContextVariable{
		{
			Default: "",
			Envar:   "MY_VARIABLE",
			Help:    "Description of my variable [%s]",
			Arg:     "my-variable",
			Type:    optiontype.String,
		},
	}

	cmdhelper.Init(cobraCommand, contextVariables)
	cmdhelper.PreRun(cobraCommand, []string{}, "example-cmd", contextVariables)
	// Output:
}

func ExampleVersion() {
	result := cmdhelper.Version("1.2.3", "4")
	fmt.Println(result)
	// Output: 1.2.3-4
}
