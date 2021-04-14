package config

const (
	defaultLogLevel = "info"
	defaultDBPath   = "./badger"
)

// Main contains main config of the App
type Main struct {
	LogLevel string `mapstructure:"log-level" json:"log-level,omitempty"`
	LogFile  string `mapstructure:"log-file" json:"log-file,omitempty"`
	Quiet    bool   `mapstructure:"quiet" json:"quiet"`
	DBPath   string `mapstructure:"db-path" json:"db-path,omitempty"`
}

// NewMain returns a new Main instance.
func NewMain() *Main {
	return &Main{
		LogLevel: defaultLogLevel,
		DBPath:   defaultDBPath,
	}
}
