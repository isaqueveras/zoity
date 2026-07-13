// Package application provides application management
package application

import (
	"zoity/domain"
	"zoity/infrastructure"
)

type Services struct {
	repositories domain.Repositories
}

// NewService creates a new instance of builder with the given repository
func NewService() *Services {
	return &Services{
		repositories: infrastructure.NewRepositories(),
	}
}
