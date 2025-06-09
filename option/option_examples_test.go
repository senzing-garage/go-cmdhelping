package option_test

import (
	"fmt"

	"github.com/senzing-garage/go-cmdhelping/option"
)

// ----------------------------------------------------------------------------
// Examples
// ----------------------------------------------------------------------------

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

func ExampleOsLookupEnvUint() {
	fmt.Println(option.OsLookupEnvUint("NOT_AN_ENVIRONMENT_VARIABLE", 10))
	// Output: 10
}

func ExampleContextVariable_SetDefault() {
	fmt.Println(option.DatabaseURL.SetDefault("NOT a default").Default)
	// Output: NOT a default
}
