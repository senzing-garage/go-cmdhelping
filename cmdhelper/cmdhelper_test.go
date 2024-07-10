package cmdhelper

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/senzing-garage/go-cmdhelping/option"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var contextVariables = []option.ContextVariable{
	option.Configuration,
	option.EngineLogLevel,
	option.EnableSwaggerUI,
	option.XtermArguments,
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestInit(test *testing.T) {
	_ = test
	cobraCommand := &cobra.Command{
		Use:   "test-use",
		Short: "test-short",
		Long:  `test-long`,
	}
	Init(cobraCommand, contextVariables)
}

func TestPreRun(test *testing.T) {
	_ = test
	cobraCommand := &cobra.Command{
		Use:   "test-use",
		Short: "test-short",
		Long:  `test-long`,
	}
	Init(cobraCommand, contextVariables)
	PreRun(cobraCommand, []string{}, "test-cmd", contextVariables)
}

func TestPreRun_badConfigurationPath(test *testing.T) {
	_ = test
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
	Init(cobraCommand, contextVariables)
	PreRun(cobraCommand, []string{}, "test-cmd", contextVariables)
}

func TestPreRun_goodConfigurationPath(test *testing.T) {
	_ = test
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
	Init(cobraCommand, contextVariables)
	PreRun(cobraCommand, []string{}, "test-cmd", contextVariables)
}

func TestVersion(test *testing.T) {
	assert.Equal(test, "1.2.3-4", Version("1.2.3", "4"))
}

func TestVersion_noIteration(test *testing.T) {
	assert.Equal(test, "1.2.3", Version("1.2.3", "0"))
}
