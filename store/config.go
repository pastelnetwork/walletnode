package store

import (
	"path"
)

type Config struct {
	path string
}

const (
	chatNamespace = "chat"
)

func NewConfig(dataDir string) *Config {
	return &Config{path: path.Join(dataDir, chatNamespace)}
}
