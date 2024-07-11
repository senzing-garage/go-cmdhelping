package option

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Tests
// ----------------------------------------------------------------------------

func TestOsLookupEnvBool(test *testing.T) {
	environmentVariableName := "AN_ENVIRONMENT_VARIABLE"
	test.Setenv(environmentVariableName, "true")
	assert.True(test, OsLookupEnvBool(environmentVariableName, false))
}

func TestOsLookupEnvBool_badValue(test *testing.T) {
	environmentVariableName := "BAD_BOOLEAN_ENVIRONMENT_VARIABLE"
	test.Setenv(environmentVariableName, "not-a-boolean")
	assert.Panics(test, func() { _ = OsLookupEnvBool(environmentVariableName, true) })
}

func TestOsLookupEnvBool_notEnvVar(test *testing.T) {
	assert.True(test, OsLookupEnvBool("NOT_AN_ENVIRONMENT_VARIABLE", true))
}

func TestOsLookupEnvBool_useDefault(test *testing.T) {
	assert.True(test, OsLookupEnvBool("NOT_AN_ENVIRONMENT_VARIABLE", true))
}

func TestOsLookupEnvInt(test *testing.T) {
	environmentVariableName := "AN_ENVIRONMENT_VARIABLE"
	test.Setenv(environmentVariableName, "10")
	assert.Equal(test, 10, OsLookupEnvInt(environmentVariableName, 99))
}

func TestOsLookupEnvInt_badValue(test *testing.T) {
	environmentVariableName := "BAD_INTEGER_ENVIRONMENT_VARIABLE"
	test.Setenv(environmentVariableName, "not-an-integer")
	assert.Panics(test, func() { _ = OsLookupEnvInt(environmentVariableName, 10) })
}

func TestOsLookupEnvInt_notEnvVar(test *testing.T) {
	assert.Equal(test, 10, OsLookupEnvInt("NOT_AN_ENVIRONMENT_VARIABLE", 10))
}

func TestOsLookupEnvInt_useDefault(test *testing.T) {
	assert.Equal(test, 10, OsLookupEnvInt("NOT_AN_ENVIRONMENT_VARIABLE", 10))
}

func TestOsLookupEnvString(test *testing.T) {
	environmentVariableName := "AN_ENVIRONMENT_VARIABLE"
	test.Setenv(environmentVariableName, "default")
	assert.Equal(test, "default", OsLookupEnvString(environmentVariableName, "not-default"))
}

func TestOsLookupEnvString_notEnvVar(test *testing.T) {
	assert.Equal(test, "default", OsLookupEnvString("NOT_AN_ENVIRONMENT_VARIABLE", "default"))
}

func TestOsLookupEnvString_useDefault(test *testing.T) {
	assert.Equal(test, "default", OsLookupEnvString("NOT_AN_ENVIRONMENT_VARIABLE", "default"))
}

func TestSetDefault(test *testing.T) {
	assert.Equal(test, "NOT a default", DatabaseURL.SetDefault("NOT a default").Default)
}
