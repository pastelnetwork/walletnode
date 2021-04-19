package dao

const (
	chatDBDir = "chat"
)

type Config struct {
	ChatDBDir string
}

// NewConfig returns a new Config instance.
func NewConfig() *Config {
	return &Config{ChatDBDir: chatDBDir}
}
