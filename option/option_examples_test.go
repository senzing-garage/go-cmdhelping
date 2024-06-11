package option

import (
	"fmt"
)

// ----------------------------------------------------------------------------
// Examples
// ----------------------------------------------------------------------------

func ExampleOsLookupEnvBool() {
	fmt.Println(OsLookupEnvBool("NOT_AN_ENVIRONMENT_VARIABLE", true))
	// Output: true
}

func ExampleOsLookupEnvInt() {
	fmt.Println(OsLookupEnvInt("NOT_AN_ENVIRONMENT_VARIABLE", 10))
	// Output: 10
}

func ExampleOsLookupEnvString() {
	fmt.Println(OsLookupEnvString("NOT_AN_ENVIRONMENT_VARIABLE", "default"))
	// Output: default
}

func ExampleContextVariable_SetDefault() {
	fmt.Println(DatabaseURL.SetDefault("NOT a default").Default)
	// Output: NOT a default
}
