// Package config provides configuration loading and management
package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

var cfg *Config

type ConfigOptions struct {
	Envs      []string
	Databases []string
	Servers   []string
}

type ConfigOption func(*ConfigOptions)

func WithDatabase(nicks ...string) ConfigOption {
	return func(co *ConfigOptions) {
		co.Databases = nicks
	}
}

func WithServer(nicks ...string) ConfigOption {
	return func(co *ConfigOptions) {
		co.Servers = nicks
	}
}

func WithEnvs(envs ...string) ConfigOption {
	return func(co *ConfigOptions) {
		co.Envs = envs
	}
}

// NewConfig returns a new config
func NewConfig(opts ...ConfigOption) (*Config, error) {
	options := ConfigOptions{Envs: []string{".env"}}
	for _, opt := range opts {
		opt(&options)
	}

	if err := godotenv.Load(options.Envs...); err != nil {
		return nil, err
	}

	cfg = new(Config)
	cfg.Databases = make(map[string]Database)
	cfg.Servers = make(map[string]Server)

	cfg.databases(options.Databases...)
	cfg.servers(options.Servers...)
	cfg.logging()

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// Config represents the application configuration
type Config struct {
	Environment string `env:"ENVIRONMENT,required" envDefault:"development"`
	App         Application
	Servers     map[string]Server
	Databases   map[string]Database
}

// IsProduction returns true if the environment is production
func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

// IsDevelopment returns true if the environment is development
func (c *Config) IsDevelopment() bool {
	return c.Environment == "development"
}

// IsTesting returns true if the environment is testing
func (c *Config) IsTesting() bool {
	return c.Environment == "testing"
}
