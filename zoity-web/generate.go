package main

//go:generate sh -c "rm -rf resources/gen/generated && mkdir -p resources/gen/generated"

//go:generate mockery

//go:generate go mod tidy --go=1.25.2

//go:generate golangci-lint run --timeout 5m --config .golangci.yml
