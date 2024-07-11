package option

import (
	"fmt"
	"time"

	"github.com/senzing-garage/go-cmdhelping/option/optiontype"
)

// Alter the default value of an option.ContextVariable.
// Useful when including an option.ContextVariable a []option.ContextVariable.
func (v ContextVariable) SetDefault(newDefault any) ContextVariable {
	v.Default = newDefault
	return v
}

var AvoidServe = ContextVariable{
	Arg:     "avoid-serving",
	Default: OsLookupEnvBool("SENZING_TOOLS_AVOID_SERVING", false),
	Envar:   "SENZING_TOOLS_AVOID_SERVING",
	Help:    "Avoid serving.  For testing only. [%s]",
	Type:    optiontype.Bool,
}

var ConfigPath = ContextVariable{
	Arg:     "config-path",
	Default: OsLookupEnvString("SENZING_TOOLS_CONFIG_PATH", ""),
	Envar:   "SENZING_TOOLS_CONFIG_PATH",
	Help:    "Path to SenzingAPI's configuration directory [%s]",
	Type:    optiontype.String,
}

var Configuration = ContextVariable{
	Arg:     "configuration",
	Default: OsLookupEnvString("SENZING_TOOLS_CONFIGURATION", ""),
	Envar:   "SENZING_TOOLS_CONFIGURATION",
	Help:    "Path to configuration file [%s]",
	Type:    optiontype.String,
}

var DatabaseURL = ContextVariable{
	Arg:     "database-url",
	Default: OsLookupEnvString("SENZING_TOOLS_DATABASE_URL", ""),
	Envar:   "SENZING_TOOLS_DATABASE_URL",
	Help:    "URL of database to initialize [%s]",
	Type:    optiontype.String,
}

var Datasources = ContextVariable{
	Arg:     "datasources",
	Default: []string{},
	Envar:   "SENZING_TOOLS_DATASOURCES",
	Help:    "Datasources to be added to initial Senzing configuration [%s]",
	Type:    optiontype.StringSlice,
}

var DelayInSeconds = ContextVariable{
	Arg:     "delay-in-seconds",
	Default: OsLookupEnvInt("SENZING_TOOLS_DELAY_IN_SECONDS", 0),
	Envar:   "SENZING_TOOLS_DELAY_IN_SECONDS",
	Help:    "Number of seconds to wait before starting process [%s]",
	Type:    optiontype.Int,
}

var EnableAll = ContextVariable{
	Arg:     "enable-all",
	Default: OsLookupEnvBool("SENZING_TOOLS_ENABLE_ALL", false),
	Envar:   "SENZING_TOOLS_ENABLE_ALL",
	Help:    "Enable all services [%s]",
	Type:    optiontype.Bool,
}

var EnableSzConfig = ContextVariable{
	Arg:     "enable-szconfig",
	Default: OsLookupEnvBool("SENZING_TOOLS_ENABLE_SZCONFIG", false),
	Envar:   "SENZING_TOOLS_ENABLE_SZCONFIG",
	Help:    "Enable SzConfig service [%s]",
	Type:    optiontype.Bool,
}

var EnableSzConfigManager = ContextVariable{
	Arg:     "enable-szconfigmanager",
	Default: OsLookupEnvBool("SENZING_TOOLS_ENABLE_SZCONFIGMANAGER", false),
	Envar:   "SENZING_TOOLS_ENABLE_SZCONFIGMANAGER",
	Help:    "Enable SZConfigManager service [%s]",
	Type:    optiontype.Bool,
}

var EnableSzDiagnostic = ContextVariable{
	Arg:     "enable-szdiagnostic",
	Default: OsLookupEnvBool("SENZING_TOOLS_ENABLE_SZDIAGNOSTIC", false),
	Envar:   "SENZING_TOOLS_ENABLE_SZDIAGNOSTIC",
	Help:    "Enable SzDiagnostic service [%s]",
	Type:    optiontype.Bool,
}

var EnableSzEngine = ContextVariable{
	Arg:     "enable-szengine",
	Default: OsLookupEnvBool("SENZING_TOOLS_ENABLE_SZENGINE", false),
	Envar:   "SENZING_TOOLS_ENABLE_SZENGINE",
	Help:    "Enable SzEngine service [%s]",
	Type:    optiontype.Bool,
}

var EnableSzProduct = ContextVariable{
	Arg:     "enable-szproduct",
	Default: OsLookupEnvBool("SENZING_TOOLS_ENABLE_SZPRODUCT", false),
	Envar:   "SENZING_TOOLS_ENABLE_SZPRODUCT",
	Help:    "Enable SZProduct service [%s]",
	Type:    optiontype.Bool,
}

var EnableSenzingChatAPI = ContextVariable{
	Arg:     "enable-senzing-chat-api",
	Default: OsLookupEnvBool("SENZING_TOOLS_ENABLE_SENZING_CHAT_API", false),
	Envar:   "SENZING_TOOLS_ENABLE_SENZING_CHAT_API",
	Help:    "Enable the Senzing REST Chat service [%s]",
	Type:    optiontype.Bool,
}

var EnableSenzingRestAPI = ContextVariable{
	Arg:     "enable-senzing-rest-api",
	Default: OsLookupEnvBool("SENZING_TOOLS_ENABLE_SENZING_REST_API", false),
	Envar:   "SENZING_TOOLS_ENABLE_SENZING_REST_API",
	Help:    "Enable the Senzing REST API service [%s]",
	Type:    optiontype.Bool,
}

var EnableSwaggerUI = ContextVariable{
	Arg:     "enable-swagger-ui",
	Default: OsLookupEnvBool("SENZING_TOOLS_ENABLE_SWAGGER_UI", false),
	Envar:   "SENZING_TOOLS_ENABLE_SWAGGER_UI",
	Help:    "Enable the Swagger UI service [%s]",
	Type:    optiontype.Bool,
}

var EnableXterm = ContextVariable{
	Arg:     "enable-xterm",
	Default: OsLookupEnvBool("SENZING_TOOLS_ENABLE_XTERM", false),
	Envar:   "SENZING_TOOLS_ENABLE_XTERM",
	Help:    "Enable the XTerm service [%s]",
	Type:    optiontype.Bool,
}

// Deprecated: Use EngineSettings instead.
var EngineConfigurationJSON = ContextVariable{
	Arg:     "engine-configuration-json",
	Default: OsLookupEnvString("SENZING_TOOLS_ENGINE_CONFIGURATION_JSON", ""),
	Envar:   "SENZING_TOOLS_ENGINE_CONFIGURATION_JSON",
	Help:    "JSON string sent to Senzing's init() function [%s]",
	Type:    optiontype.String,
}

var EngineInstanceName = ContextVariable{
	Arg:     "engine-instance-name",
	Default: fmt.Sprintf("senzing-tools-%d", time.Now().Unix()),
	Envar:   "SENZING_TOOLS_ENGINE_INSTANCE_NAME",
	Help:    "Identifier given to the Senzing engine [%s]",
	Type:    optiontype.String,
}

var EngineLogLevel = ContextVariable{
	Arg:     "engine-log-level",
	Default: OsLookupEnvInt("SENZING_TOOLS_ENGINE_LOG_LEVEL", 0),
	Envar:   "SENZING_TOOLS_ENGINE_LOG_LEVEL",
	Help:    "Log level for Senzing Engine [%s]",
	Type:    optiontype.Int,
}

// Deprecated: Use EngineInstanceName instead.
var EngineModuleName = ContextVariable{
	Arg:     "engine-module-name",
	Default: fmt.Sprintf("senzing-tools-%d", time.Now().Unix()),
	Envar:   "SENZING_TOOLS_ENGINE_MODULE_NAME",
	Help:    "Identifier given to the Senzing engine [%s]",
	Type:    optiontype.String,
}

var EngineSettings = ContextVariable{
	Arg:     "engine-settings",
	Default: OsLookupEnvString("SENZING_TOOLS_ENGINE_SETTINGS", ""),
	Envar:   "SENZING_TOOLS_ENGINE_SETTINGS",
	Help:    "JSON string sent to Senzing's init() function [%s]",
	Type:    optiontype.String,
}

var MessageID = ContextVariable{
	Arg:     "message-id",
	Default: OsLookupEnvString("SENZING_TOOLS_MESSAGE_ID", ""),
	Envar:   "SENZING_TOOLS_MESSAGE_ID",
	Help:    "Give an explanation of a specific Senzing message [%s]",
	Type:    optiontype.String,
}

var GrpcPort = ContextVariable{
	Arg:     "grpc-port",
	Default: OsLookupEnvInt("SENZING_TOOLS_GRPC_PORT", 8261),
	Envar:   "SENZING_TOOLS_GRPC_PORT",
	Help:    "Port used to serve gRPC [%s]",
	Type:    optiontype.Int,
}

var GrpcURL = ContextVariable{
	Arg:     "grpc-url",
	Default: OsLookupEnvString("SENZING_TOOLS_GRPC_URL", ""),
	Envar:   "SENZING_TOOLS_GRPC_URL",
	Help:    "URL of Senzing gRPC service [%s]",
	Type:    optiontype.String,
}

var HTTPPort = ContextVariable{
	Arg:     "http-port",
	Default: OsLookupEnvInt("SENZING_TOOLS_HTTP_PORT", 8260),
	Envar:   "SENZING_TOOLS_HTTP_PORT",
	Help:    "Port to serve HTTP [%s]",
	Type:    optiontype.Int,
}

var InputFileType = ContextVariable{
	Arg:     "input-file-type",
	Default: OsLookupEnvString("SENZING_TOOLS_INPUT_FILE_TYPE", ""),
	Envar:   "SENZING_TOOLS_INPUT_FILE_TYPE",
	Help:    "Input file type to override auto-detect based on file name [%s]",
	Type:    optiontype.String,
}

var InputURL = ContextVariable{
	Arg:     "input-url",
	Default: OsLookupEnvString("SENZING_TOOLS_INPUT_URL", ""),
	Envar:   "SENZING_TOOLS_INPUT_URL",
	Help:    "Input URL used for processing [%s]",
	Type:    optiontype.String,
}

var JSONOutput = ContextVariable{
	Arg:     "json-output",
	Default: OsLookupEnvBool("SENZING_TOOLS_JSON_OUTPUT", false),
	Envar:   "SENZING_TOOLS_JSON_OUTPUT",
	Help:    "Only output JSON messages  [%s]",
	Type:    optiontype.Bool,
}

var LicenseDaysLeft = ContextVariable{
	Arg:     "license-days-left",
	Default: OsLookupEnvString("SENZING_TOOLS_LICENSE_DAYS_LEFT", "30"),
	Envar:   "SENZING_TOOLS_LICENSE_DAYS_LEFT",
	Help:    "Number of days left in license before flagging as an error [%s]",
	Type:    optiontype.String,
}

var LicenseRecordsPercent = ContextVariable{
	Arg:     "license-records-percent",
	Default: OsLookupEnvString("SENZING_TOOLS_LICENSE_RECORDS_PERCENT", "90"),
	Envar:   "SENZING_TOOLS_LICENSE_RECORDS_PERCENT",
	Help:    "Percent of records allowed by license [%s]",
	Type:    optiontype.String,
}

var LicenseStringBase64 = ContextVariable{
	Arg:     "license-string-base64",
	Default: OsLookupEnvString("SENZING_TOOLS_LICENSE_STRING_BASE64", ""),
	Envar:   "SENZING_TOOLS_LICENSE_STRING_BASE64",
	Help:    "Base64 representation of a Senzing license [%s]",
	Type:    optiontype.String,
}

var LogLevel = ContextVariable{
	Arg:     "log-level",
	Default: OsLookupEnvString("SENZING_TOOLS_LOG_LEVEL", "INFO"),
	Envar:   "SENZING_TOOLS_LOG_LEVEL",
	Help:    "Log level of TRACE, DEBUG, INFO, WARN, ERROR, FATAL, or PANIC [%s]",
	Type:    optiontype.String,
}

var MonitoringPeriodInSeconds = ContextVariable{
	Arg:     "monitoring-period-in-seconds",
	Default: OsLookupEnvInt("SENZING_TOOLS_MONITORING_PERIOD_IN_SECONDS", 600),
	Envar:   "SENZING_TOOLS_MONITORING_PERIOD_IN_SECONDS",
	Help:    "Print monitoring log messages with the period given in seconds [%s]",
	Type:    optiontype.Int,
}

var NumberOfWorkers = ContextVariable{
	Arg:     "number-of-workers",
	Default: OsLookupEnvInt("SENZING_TOOLS_NUMBER_OF_WORKERS", 0),
	Envar:   "SENZING_TOOLS_NUMBER_OF_WORKERS",
	Help:    "Override the default number of worker routines. Default is GOMAXPROCS [%s]",
	Type:    optiontype.Int,
}

var ObserverOrigin = ContextVariable{
	Arg:     "observer-origin",
	Default: OsLookupEnvString("SENZING_TOOLS_OBSERVER_ORIGIN", ""),
	Envar:   "SENZING_TOOLS_OBSERVER_ORIGIN",
	Help:    "Identify this instance to the Observer [%s]",
	Type:    optiontype.String,
}

var ObserverGrpcPort = ContextVariable{
	Arg:     "observer-grpc-port",
	Default: OsLookupEnvInt("SENZING_TOOLS_OBSERVER_GRPC_PORT", 8258),
	Envar:   "SENZING_TOOLS_OBSERVER_GRPC_PORT",
	Help:    "Port to serve gRPC for receiving observer messages [%s]",
	Type:    optiontype.Int,
}

var ObserverURL = ContextVariable{
	Arg:     "observer-url",
	Default: OsLookupEnvString("SENZING_TOOLS_OBSERVER_URL", ""),
	Envar:   "SENZING_TOOLS_OBSERVER_URL",
	Help:    "URL of Observer [%s]",
	Type:    optiontype.String,
}

var OutputURL = ContextVariable{
	Arg:     "output-url",
	Default: OsLookupEnvString("SENZING_TOOLS_OUTPUT_URL", ""),
	Envar:   "SENZING_TOOLS_OUTPUT_URL",
	Help:    "Output URL used for processing [%s]",
	Type:    optiontype.String,
}

var RecordMax = ContextVariable{
	Arg:     "record-max",
	Default: OsLookupEnvInt("SENZING_TOOLS_RECORD_MAX", 0),
	Envar:   "SENZING_TOOLS_RECORD_MAX",
	Help:    "Process a maximum number of records equal to this number [%s]",
	Type:    optiontype.Int,
}

var RecordMin = ContextVariable{
	Arg:     "record-min",
	Default: OsLookupEnvInt("SENZING_TOOLS_RECORD_MIN", 0),
	Envar:   "SENZING_TOOLS_RECORD_MIN",
	Help:    "Process records starting at this record number, discarding all before [%s]",
	Type:    optiontype.Int,
}

var RecordMonitor = ContextVariable{
	Arg:     "record-monitor",
	Default: OsLookupEnvInt("SENZING_TOOLS_RECORD_MONITOR", 100000),
	Envar:   "SENZING_TOOLS_RECORD_MONITOR",
	Help:    "Log a monitor message after this number of records have been processed [%s]",
	Type:    optiontype.Int,
}

var ResourcePath = ContextVariable{
	Arg:     "resource-path",
	Default: OsLookupEnvString("SENZING_TOOLS_RESOURCE_PATH", ""),
	Envar:   "SENZING_TOOLS_RESOURCE_PATH",
	Help:    "Path to SenzingAPI's config, schema, and templates directory [%s]",
	Type:    optiontype.String,
}

var SenzingDirectory = ContextVariable{
	Arg:     "senzing-directory",
	Default: OsLookupEnvString("SENZING_TOOLS_SENZING_DIRECTORY", ""),
	Envar:   "SENZING_TOOLS_SENZING_DIRECTORY",
	Help:    "Path to the SenzingAPI installation directory [%s]",
	Type:    optiontype.String,
}

var ServerAddress = ContextVariable{
	Arg:     "server-address",
	Default: OsLookupEnvString("SENZING_TOOLS_SERVER_ADDRESS", "0.0.0.0"),
	Envar:   "SENZING_TOOLS_SERVER_ADDRESS",
	Help:    "IP interface server listens on [%s]",
	Type:    optiontype.String,
}

var SupportPath = ContextVariable{
	Arg:     "support-path",
	Default: OsLookupEnvString("SENZING_TOOLS_SUPPORT_PATH", ""),
	Envar:   "SENZING_TOOLS_SUPPORT_PATH",
	Help:    "Path to SenzingAPI's data directory [%s]",
	Type:    optiontype.String,
}

var TtyOnly = ContextVariable{
	Arg:     "tty-only",
	Default: OsLookupEnvBool("SENZING_TOOLS_TTY_ONLY", false),
	Envar:   "SENZING_TOOLS_TTY_ONLY",
	Help:    "Output confined to terminal (TTY) [%s]",
	Type:    optiontype.Bool,
}

var VisibilityPeriodInSeconds = ContextVariable{
	Arg:     "visibility-period-in-seconds",
	Default: OsLookupEnvInt("SENZING_TOOLS_VISIBILITY_PERIOD_IN_SECONDS", 60),
	Envar:   "SENZING_TOOLS_VISIBILITY_PERIOD_IN_SECONDS",
	Help:    "Number of seconds a record held for processing.  This is renewed if processing takes longer [%s]",
	Type:    optiontype.Int,
}

var XtermAllowedHostnames = ContextVariable{
	Arg:     "xterm-allowed-hostnames",
	Default: []string{"localhost"},
	Envar:   "SENZING_TOOLS_XTERM_ALLOWED_HOSTNAMES",
	Help:    "Comma-delimited list of hostnames permitted to connect to the websocket [%s]",
	Type:    optiontype.StringSlice,
}

var XtermArguments = ContextVariable{
	Arg:     "xterm-arguments",
	Default: []string{},
	Envar:   "SENZING_TOOLS_XTERM_ARGUMENTS",
	Help:    "Comma-delimited list of arguments passed to the terminal command prompt [%s]",
	Type:    optiontype.StringSlice,
}

var XtermCommand = ContextVariable{
	Arg:     "xterm-command",
	Default: OsLookupEnvString("SENZING_TOOLS_XTERM_COMMAND", "/bin/bash"),
	Envar:   "SENZING_TOOLS_XTERM_COMMAND",
	Help:    "Path of shell command [%s]",
	Type:    optiontype.String,
}

var XtermConnectionErrorLimit = ContextVariable{
	Arg:     "xterm-connection-error-limit",
	Default: OsLookupEnvInt("SENZING_TOOLS_XTERM_CONNECTION_ERROR_LIMIT", 10),
	Envar:   "SENZING_TOOLS_XTERM_CONNECTION_ERROR_LIMIT",
	Help:    "Connection re-attempts before terminating [%s]",
	Type:    optiontype.Int,
}

var XtermKeepalivePingTimeout = ContextVariable{
	Arg:     "xterm-keepalive-ping-timeout",
	Default: OsLookupEnvInt("SENZING_TOOLS_XTERM_KEEPALIVE_PING_TIMEOUT", 20),
	Envar:   "SENZING_TOOLS_XTERM_KEEPALIVE_PING_TIMEOUT",
	Help:    "Maximum allowable seconds between a ping message and its response [%s]",
	Type:    optiontype.Int,
}

var XtermMaxBufferSizeBytes = ContextVariable{
	Arg:     "xterm-max-buffer-size-bytes",
	Default: OsLookupEnvInt("SENZING_TOOLS_XTERM_MAX_BUFFER_SIZE_BYTES", 512),
	Envar:   "SENZING_TOOLS_XTERM_MAX_BUFFER_SIZE_BYTES",
	Help:    "Maximum length of terminal input [%s]",
	Type:    optiontype.Int,
}
