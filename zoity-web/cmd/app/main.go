// Package main contains the main function of the app
package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"zoity/core"
	"zoity/core/config"
	"zoity/resources/pages"

	"github.com/gin-gonic/gin"
)

func main() {
	_, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	app := core.NewApp(cfg)

	app.Router.GET("/dashboard", func(ctx *gin.Context) {
		pages.DashboardPage().Render(ctx.Request.Context(), ctx.Writer)
	})

	app.Router.GET("/services", func(ctx *gin.Context) {
		pages.ListServicesPage().Render(ctx.Request.Context(), ctx.Writer)
	})

	app.Router.GET("/flows", func(ctx *gin.Context) {
		pages.ListFlowsPage().Render(ctx.Request.Context(), ctx.Writer)
	})

	app.Router.GET("/settings", func(ctx *gin.Context) {
		pages.ListSettingsPage().Render(ctx.Request.Context(), ctx.Writer)
	})

	if err := app.Router.Run(app.Config.Servers[config.ServerDefault].GetAddress()); err != nil {
		slog.Error("Erro ao inicializar servidor da aplicação")
	}
}
