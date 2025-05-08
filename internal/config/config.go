package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// Config holds all configuration for the WAL Guardian
type Config struct {
	WALPath     string
	MaxSize     int64
	MaxAge      int64
	DryRun      bool
	LogLevel    string
}

// NewConfig creates a new Config instance from command line flags
func NewConfig() (*Config, error) {
	cfg := &Config{}

	flag.StringVar(&cfg.WALPath, "path", "", "Path to Prometheus WAL directory")
	flag.Int64Var(&cfg.MaxSize, "max-size", 5*1024*1024*1024, "Maximum WAL size in bytes (default: 5GB)")
	flag.Int64Var(&cfg.MaxAge, "max-age", 0, "Maximum age of WAL segments in hours (0 to disable)")
	flag.BoolVar(&cfg.DryRun, "dry-run", false, "Only check without making changes")
	flag.StringVar(&cfg.LogLevel, "log-level", "info", "Log level (debug, info, warn, error)")

	flag.Parse()

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

// validate checks if the configuration is valid
func (c *Config) validate() error {
	if c.WALPath == "" {
		return fmt.Errorf("WAL path is required")
	}

	// Check if WAL directory exists
	if _, err := os.Stat(c.WALPath); os.IsNotExist(err) {
		return fmt.Errorf("WAL directory does not exist: %s", c.WALPath)
	}

	// Ensure WAL path is absolute
	absPath, err := filepath.Abs(c.WALPath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %v", err)
	}
	c.WALPath = absPath

	return nil
} 