package config

// Application represents the application settings
type Application struct {
	Name    string `env:"APP_NAME" envDefault:"zoity"`
	Version string `env:"APP_VERSION" envDefault:"v0.0.0"`
	Debug   bool   `env:"APP_DEBUG" envDefault:"false"`
}
