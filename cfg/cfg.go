package cfg

import (
	"log"

	"github.com/spf13/viper"
)

type Project struct {
	Include []string
	Exclude []string
}

type Bucket struct {
	Include []string
	Exclude []string
}

type Provider struct {
	Name     string
	Enabled  bool
	Projects []Project
	Buckets  []Bucket
}

type LogStream struct {
	Enabled   bool
	Log_level string
	Format    string
}

type LogFile struct {
	Enabled   bool
	Log_level string
	Format    string
	Name      string
	Path      string
}

type LoggingConfig struct {
	LogStream LogStream
	LogFile   LogFile
}

type ReportingConfig struct {
	Enabled bool
	Format  string
	Name    string
	Path    string
}

type Config struct {
	Providers []Provider
	Logging   []LoggingConfig
	Reporting ReportingConfig
}

type ConfigRepository interface {
	SetDefaultValues()
	Load() (*Config, error)
}

func (config *Config) SetDefaultValues() {
	if config.Providers == nil {
		config.Providers = []Provider{
			{
				Name:     "gcp",
				Enabled:  true,
				Projects: []Project{},
				Buckets:  []Bucket{},
			},
		}
	}

	if config.Logging == nil {
		config.Logging = []LoggingConfig{
			{
				LogStream: LogStream{
					Enabled:   true,
					Log_level: DEFAULT_LOG_LEVEL,
					Format:    DEFAULT_LOG_FORMAT,
				},
			},
		}
	}

	if !config.Reporting.Enabled {
		config.Reporting = ReportingConfig{
			Enabled: true,
			Format:  DEFAULT_REPORT_FILE_EXTENSION,
			Name:    DEFAULT_REPORT_FILE_NAME,
			Path:    DEFAULT_REPORT_FILE_PATH,
		}
	}
}

func Load() (*Config, error) {
	viper.SetConfigName("config.yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fail to read config: %v", err)
	}

	var config_yaml Config
	if err := viper.Unmarshal(&config_yaml); err != nil {
		log.Fatalf("Fail to parse config: %v", err)
	}

	return &config_yaml, nil
}
