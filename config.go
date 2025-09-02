package nftables

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
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

func (c *Config) do(args []string) error {
	stdout, stderr, err := c.exec(args)
	if err != nil {
		return err
	}

	if stderr != "" {
		return errors.New(stderr)
	}

	if stdout != "" {
		return errors.New(stdout)
	}

	return nil
}

func (c *Config) exec(args []string) (string, string, error) {
	c.Logger.Printf("exec '%s %s'", c.Path, strings.Join(args, " "))
	cmd := exec.CommandContext(context.Background(), c.Path, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		c.Logger.Printf("exec error '%s' '%s' '%s'", stdout.String(), stderr.String(), err.Error())
		return stdout.String(), stderr.String(), fmt.Errorf("error while executing command: %v", err)
	}

	return stdout.String(), stderr.String(), nil
}
