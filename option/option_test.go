package option_test

import (
	"testing"

	"github.com/senzing-garage/go-cmdhelping/option"
	"github.com/stretchr/testify/assert"
)

const genericEnvVarName = "AN_ENVIRONMENT_VARIABLE"

// ----------------------------------------------------------------------------
// Tests
// ----------------------------------------------------------------------------

func TestOsLookupEnvBool(test *testing.T) {
	test.Setenv(genericEnvVarName, "true")
	assert.True(test, option.OsLookupEnvBool(genericEnvVarName, false))
}

func TestOsLookupEnvBool_badValue(test *testing.T) {
	environmentVariableName := "BAD_BOOLEAN_ENVIRONMENT_VARIABLE"
	test.Setenv(environmentVariableName, "not-a-boolean")
	assert.Panics(test, func() { _ = option.OsLookupEnvBool(environmentVariableName, true) })
}

func TestOsLookupEnvBool_notEnvVar(test *testing.T) {
	test.Parallel()
	assert.True(test, option.OsLookupEnvBool("NOT_AN_ENVIRONMENT_VARIABLE", true))
}

func TestOsLookupEnvBool_useDefault(test *testing.T) {
	test.Parallel()
	assert.True(test, option.OsLookupEnvBool("NOT_AN_ENVIRONMENT_VARIABLE", true))
}

func TestOsLookupEnvInt(test *testing.T) {
	test.Setenv(genericEnvVarName, "10")
	assert.Equal(test, 10, option.OsLookupEnvInt(genericEnvVarName, 99))
}

func TestOsLookupEnvInt_badValue(test *testing.T) {
	environmentVariableName := "BAD_INTEGER_ENVIRONMENT_VARIABLE"
	test.Setenv(environmentVariableName, "not-an-integer")
	assert.Panics(test, func() { _ = option.OsLookupEnvInt(environmentVariableName, 10) })
}

func TestOsLookupEnvInt_notEnvVar(test *testing.T) {
	test.Parallel()
	assert.Equal(test, 10, option.OsLookupEnvInt("NOT_AN_ENVIRONMENT_VARIABLE", 10))
}

func TestOsLookupEnvInt_useDefault(test *testing.T) {
	test.Parallel()
	assert.Equal(test, 10, option.OsLookupEnvInt("NOT_AN_ENVIRONMENT_VARIABLE", 10))
}

func TestOsLookupEnvString(test *testing.T) {
	test.Setenv(genericEnvVarName, "default")
	assert.Equal(test, "default", option.OsLookupEnvString(genericEnvVarName, "not-default"))
}

func TestOsLookupEnvString_notEnvVar(test *testing.T) {
	test.Parallel()
	assert.Equal(test, "default", option.OsLookupEnvString("NOT_AN_ENVIRONMENT_VARIABLE", "default"))
}

func TestOsLookupEnvString_useDefault(test *testing.T) {
	test.Parallel()
	assert.Equal(test, "default", option.OsLookupEnvString("NOT_AN_ENVIRONMENT_VARIABLE", "default"))
}

func TestSetDefault(test *testing.T) {
	test.Parallel()
	assert.Equal(test, "NOT a default", option.DatabaseURL.SetDefault("NOT a default").Default)
}
