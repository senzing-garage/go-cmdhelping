package cmdhelper_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/senzing-garage/go-cmdhelping/cmdhelper"
	"github.com/senzing-garage/go-cmdhelping/option"
	"github.com/senzing-garage/go-cmdhelping/option/optiontype"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var uintOption = option.ContextVariable{
	Arg:     "uint-option",
	Default: option.OsLookupEnvUint("SENZING_TOOLS_TEST_UINT_OPTION", 10),
	Envar:   "SENZING_TOOLS_TEST_UINT_OPTION",
	Help:    "Test help. [%s]",
	Type:    optiontype.Uint,
}

var uint32Option = option.ContextVariable{
	Arg:     "uint-32-option",
	Default: option.OsLookupEnvUint32("SENZING_TOOLS_TEST_UINT_32_OPTION", 10),
	Envar:   "SENZING_TOOLS_TEST_UINT_32_OPTION",
	Help:    "Test help. [%s]",
	Type:    optiontype.Uint32,
}

var uint64Option = option.ContextVariable{
	Arg:     "uint-64-option",
	Default: option.OsLookupEnvUint64("SENZING_TOOLS_TEST_UINT_64_OPTION", 10),
	Envar:   "SENZING_TOOLS_TEST_UINT_64_OPTION",
	Help:    "Test help. [%s]",
	Type:    optiontype.Uint64,
}

var contextVariables = []option.ContextVariable{
	option.Configuration,
	option.EnableSwaggerUI,
	option.CoreLogLevel,
	option.XtermArguments,
	uint32Option,
	uint64Option,
	uintOption,
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestInit(test *testing.T) {
	test.Parallel()

	cobraCommand := &cobra.Command{
		Use:   "test-use",
		Short: "test-short",
		Long:  `test-long`,
	}
	cmdhelper.Init(cobraCommand, contextVariables)
}

func TestPreRun(test *testing.T) {
	test.Parallel()

	cobraCommand := &cobra.Command{
		Use:   "test-use",
		Short: "test-short",
		Long:  `test-long`,
	}
	cmdhelper.Init(cobraCommand, contextVariables)
	cmdhelper.PreRun(cobraCommand, []string{}, "test-cmd", contextVariables)
}

func TestPreRun_badConfigurationPath(test *testing.T) {
	test.Parallel()

	configurationOption := option.Configuration
	configurationOption.Default = "/tmp/senzing-tools"
	contextVariables := []option.ContextVariable{
		configurationOption,
	}
	cobraCommand := &cobra.Command{
		Use:   "test-use",
		Short: "test-short",
		Long:  `test-long`,
	}
	cmdhelper.Init(cobraCommand, contextVariables)
	cmdhelper.PreRun(cobraCommand, []string{}, "test-cmd", contextVariables)
}

func TestPreRun_goodConfigurationPath(test *testing.T) {
	test.Parallel()
	configurationFilePath := filepath.Join(test.TempDir(), "configuration.txt")
	file, err := os.Create(configurationFilePath)
	require.NoError(test, err)

	defer file.Close()

	configurationOption := option.Configuration
	configurationOption.Default = configurationFilePath
	contextVariables := []option.ContextVariable{
		configurationOption,
	}
	cobraCommand := &cobra.Command{
		Use:   "test-use",
		Short: "test-short",
		Long:  `test-long`,
	}
	cmdhelper.Init(cobraCommand, contextVariables)
	cmdhelper.PreRun(cobraCommand, []string{}, "test-cmd", contextVariables)
}

func TestVersion(test *testing.T) {
	test.Parallel()
	assert.Equal(test, "1.2.3-4", cmdhelper.Version("1.2.3", "4"))
}

func TestVersion_noIteration(test *testing.T) {
	test.Parallel()
	assert.Equal(test, "1.2.3", cmdhelper.Version("1.2.3", "0"))
}
