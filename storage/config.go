package storage

type Config struct {
	ChatDBDir string
}

const (
	chatNamespace = "chat"
)

func NewConfig() *Config {
	return &Config{ChatDBDir: chatNamespace}
}
