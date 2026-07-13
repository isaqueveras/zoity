// Package main contains the main function of the app
package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"zoity/core"
	"zoity/core/config"
	"zoity/services/database"
)

func main() {
	_, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	app := core.NewApp(cfg)

	conn := database.OpenConnections(app)
	defer conn.CloseConnections()

	if err := app.Router.Run(app.Config.Servers[config.ServerDefault].GetAddress()); err != nil {
		slog.Error("Erro ao inicializar servidor da aplicação")
	}
}
