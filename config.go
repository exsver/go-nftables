package nftables

import (
	"io"
	"log"
)

type Config struct {
	// Path to nft bin
	// Default:
	//   "/usr/sbin/nft"
	Path string
	// Chain
	Chain *Chain
	// Logger - debug logger
	Logger *log.Logger
}

func NewConfig(chain *Chain) (*Config, error) {
	return &Config{
		Path:   "/usr/sbin/nft",
		Chain:  chain,
		Logger: log.New(io.Discard, "", 0),
	}, nil
}

func (c *Config) SetLogger(logger *log.Logger) {
	c.Logger = logger
}

func (c *Config) SetPath(path string) {
	c.Path = path
}
