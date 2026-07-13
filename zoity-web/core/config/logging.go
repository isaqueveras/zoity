package config

import (
	"log/slog"
	"os"
)

func (c *Config) logging() {
	var loggHandler slog.Handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})
	if c.IsDevelopment() {
		loggHandler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	}
	slog.SetDefault(slog.New(loggHandler))
}
