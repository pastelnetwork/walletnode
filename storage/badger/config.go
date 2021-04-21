package badger

// Config contains settings of the badger-database.
type Config struct {
	ChatDBDir string `mapstructure:"badger-dir" json:"badger-dir,omitempty"`
}

// NewConfig returns a new Config instance.
func NewConfig() *Config {
	return &Config{}
}
