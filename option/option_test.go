package option

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Tests
// ----------------------------------------------------------------------------

func TestOsLookupEnvBool(test *testing.T) {
	test.Setenv("AN_ENVIRONMENT_VARIABLE", "true")
	assert.True(test, OsLookupEnvBool("AN_ENVIRONMENT_VARIABLE", false))
}

func TestOsLookupEnvBool_useDefault(test *testing.T) {
	assert.True(test, OsLookupEnvBool("NOT_AN_ENVIRONMENT_VARIABLE", true))
}

func TestOsLookupEnvInt(test *testing.T) {
	test.Setenv("AN_ENVIRONMENT_VARIABLE", "10")
	assert.Equal(test, 10, OsLookupEnvInt("AN_ENVIRONMENT_VARIABLE", 99))
}

func TestOsLookupEnvInt_useDefault(test *testing.T) {
	assert.Equal(test, 10, OsLookupEnvInt("NOT_AN_ENVIRONMENT_VARIABLE", 10))
}

func TestOsLookupEnvString(test *testing.T) {
	test.Setenv("AN_ENVIRONMENT_VARIABLE", "default")
	assert.Equal(test, "default", OsLookupEnvString("AN_ENVIRONMENT_VARIABLE", "not-default"))
}

func TestOsLookupEnvString_useDefault(test *testing.T) {
	assert.Equal(test, "default", OsLookupEnvString("NOT_AN_ENVIRONMENT_VARIABLE", "default"))
}

func TestSetDefault(test *testing.T) {
	assert.Equal(test, "NOT a default", DatabaseURL.SetDefault("NOT a default").Default)
}
