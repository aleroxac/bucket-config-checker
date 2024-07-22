package cfg

const (
	DEFAULT_CONFIG_FILE_NAME      = "config"
	DEFAULT_CONFIG_FILE_EXTENSION = "yaml"
	DEFAULT_CONFIG_FILE_PATH      = "."
)

const (
	DEFAULT_LOG_LEVEL  = "info"
	DEFAULT_LOG_FORMAT = "tty"
)

const (
	DEFAULT_LOG_FILE_NAME      = "report"
	DEFAULT_LOG_FILE_EXTENSION = "yaml"
	DEFAULT_LOG_FILE_PATH      = "."
)

const (
	DEFAULT_REPORT_FILE_NAME      = "report"
	DEFAULT_REPORT_FILE_EXTENSION = "yaml"
	DEFAULT_REPORT_FILE_PATH      = "."
)

var ALLOWED_CONFIG_EXTENSIONS = []string{
	"yaml",
	"json",
	"env",
}

var ALLOWED_REPORT_EXTENSIONS = []string{
	"yaml",
	"json",
	"csv",
	"svg",
	"html",
	"table",
}
