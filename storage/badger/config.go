package badger

// Config contains settings of the badger-database.
type Config struct {
	Dir string `mapstructure:"db-dir" json:"db-dir,omitempty"`
}

// NewConfig returns a new Config instance.
func NewConfig() *Config {
	return &Config{}
}
