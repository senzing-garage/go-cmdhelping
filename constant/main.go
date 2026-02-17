package constant //revive:disable-line var-naming

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

const (
	SetEnvPrefix    = "SENZING_TOOLS"
	VersionTemplate = `{{printf "%s: %s - version %s\n" .Name .Short .Version}}`
)
