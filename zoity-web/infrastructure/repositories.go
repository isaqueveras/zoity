// Package infrastructure provides infrastructure services and builders
package infrastructure

type repositories struct{}

// NewRepositories creates a new instance of builder for infrastructure
func NewRepositories() *repositories {
	return &repositories{}
}
