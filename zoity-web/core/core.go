package core

import (
	"zoity/core/config"
	"zoity/helpers"
	"zoity/resources/layouts"

	"github.com/gin-gonic/gin"
)

// App defines the core framework
type App struct {
	Config *config.Config
	Router *gin.Engine
}

// NewApp returns a new app instance
func NewApp(cfg *config.Config) *App {
	r := gin.New()

	r.NoRoute(func(ctx *gin.Context) {
		helpers.Render(ctx, layouts.NotFound())
	})

	return &App{Config: cfg, Router: r}
}
