package config

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/caarlos0/env/v11"
)

const (
	// ServerDefault define o nome da aplicação principal
	ServerDefault string = "default"
)

// Server define a estrutura de dados para configurar um servidor
type Server struct {
	Host string `env:"SERVER_HOST" envDefault:"localhost"`
	Port string `env:"SERVER_PORT" envDefault:"8080"`
}

// GetAddress obtem a url completa de um servidor
func (s Server) GetAddress() string {
	return s.Host + ":" + s.Port
}

func (c *Config) servers(nicks ...string) {
	if len(nicks) == 0 {
		slog.Warn("Nenhum servidor configurado, usando o default")
		c.servers(ServerDefault)
		return
	}

	for _, name := range nicks {
		server := &Server{}
		opts := env.Options{Prefix: fmt.Sprintf("%s_", strings.ToUpper(name))}
		if err := env.ParseWithOptions(server, opts); err != nil {
			slog.Error("Error loading server configuration", slog.String("server", name), slog.String("error", err.Error()))
			continue
		}
		c.Servers[name] = *server
		slog.Info("Server configured", slog.String("server", name))
	}
}
