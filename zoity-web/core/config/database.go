package config

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/caarlos0/env/v11"
)

const (
	DatabaseDefault string = "default"
)

// Database represents the database connection settings
type Database struct {
	Nick      string `env:"DATABASE_NICK" envDefault:"default"`
	Name      string `env:"DATABASE_NAME"`
	Username  string `env:"DATABASE_USER" envDefault:"postgres"`
	Password  string `env:"DATABASE_PASS" envDefault:"postgres"`
	Host      string `env:"DATABASE_HOST" envDefault:"localhost"`
	Port      string `env:"DATABASE_PORT" envDefault:"5432"`
	MaxConn   int    `env:"DATABASE_MAX_CONN" envDefault:"10"`
	MaxIdle   int    `env:"DATABASE_MAX_IDLE" envDefault:"5"`
	ReadOnly  bool   `env:"DATABASE_READ_ONLY" envDefault:"false"`
	Timeout   int32  `env:"DATABASE_TIMEOUT" envDefault:"60"`
	SSLMode   string `env:"DATABASE_SSL_MODE" envDefault:"disable"`
	SSLClient SSLCertificate
}

// SSLCertificate represents the SSL certificate settings
type SSLCertificate struct {
	Certificate          string `env:"DATABASE_SSL_CERT"`
	PrivateKey           string `env:"DATABASE_SSL_KEY"`
	CertificateAuthority string `env:"DATABASE_SSL_CA"`
}

// GetDatabase return the database config
func (c *Config) GetDatabase(nick string) *Database {
	for idx := range c.Databases {
		if c.Databases[idx].Nick != "" && c.Databases[idx].Nick == nick {
			return c.GetDatabase(c.Databases[idx].Nick)
		}
	}
	return nil
}

func (c *Config) databases(nicks ...string) {
	if len(nicks) == 0 {
		slog.Warn("Nenhum banco de dados configurado, usando o default")
		c.databases(DatabaseDefault)
		return
	}

	for _, name := range nicks {
		db := &Database{}
		opts := env.Options{Prefix: fmt.Sprintf("%s_", strings.ToUpper(name))}
		if err := env.ParseWithOptions(db, opts); err != nil {
			slog.Error("Error loading database configuration", slog.String("database", name), slog.String("error", err.Error()))
			continue
		}
		c.Databases[name] = *db
		slog.Info("Database configured", slog.String("database", name))
	}
}
