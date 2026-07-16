dev:
	templ generate -watch -cmd "go run cmd/app/main.go"

gogen:
	go generate ./...

air: clean
	air -c .air.toml

clean:
	@echo "> Cleaning project..."
	@go clean
	@rm -rf ./bin

open-test:
	xdg-open coverage.html

test:
	@echo "> Running tests..."
	@go test -coverprofile=coverage.out ./...
	@sed -i '/_templ\.go/d' coverage.out
	@sed -i '/mock/d' coverage.out
	@go tool cover -html=coverage.out -o coverage.html

templ: 
	@go run github.com/a-h/templ/cmd/templ@latest generate

build: clean templ gogen test
	@echo "> Building project..."
	go build -o ./bin/app ./cmd/app/main.go

run-build: build
	./bin/app

lint:
	@echo "> Linting project..."
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

.PHONY: dev clean test lint build run-build
