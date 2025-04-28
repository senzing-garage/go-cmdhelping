package option

import (
	"os"
	"strconv"
	"strings"
)

// If the environment variable exists, return its value.  If not, return the default value.
func OsLookupEnvBool(envar string, aDefault bool) bool {
	resultString, isSet := os.LookupEnv(envar)
	if !isSet {
		return aDefault
	}

	result, err := strconv.ParseBool(strings.ToUpper(resultString))
	if err != nil {
		panic(err)
	}

	return result
}

// If the environment variable exists, return its value.  If not, return the default value.
func OsLookupEnvInt(envar string, aDefault int) int {
	resultString, isSet := os.LookupEnv(envar)
	if !isSet {
		return aDefault
	}

	result, err := strconv.Atoi(resultString)
	if err != nil {
		panic(err)
	}

	return result
}

// If the environment variable exists, return its value.  If not, return the default value.
func OsLookupEnvString(envar string, aDefault string) string {
	result, isSet := os.LookupEnv(envar)
	if !isSet {
		return aDefault
	}

	return result
}
