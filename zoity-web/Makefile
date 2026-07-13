dev:
	templ generate -watch -cmd "go run cmd/app/main.go"

gogen:
	go generate ./...

gen:
	templ generate && go run cmd/generator/main.go

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
	@go run github.com/a-h/templ/cmd/templ@latest generate
	go build -o ./bin/app ./cmd/app/main.go

run-build: build
	./bin/app

lint:
	@echo "> Linting project..."
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

.PHONY: dev clean test lint build run-build

# --------------- Docker Tools ---------------

DOCKER_COMMAND :=	docker compose -p zoity -f docker-compose.yml

docker-dev: 
	@echo "> 🐳 Run development server"
	@$(DOCKER_COMMAND) up zoity_postgres zoity_app

docker-app:
	@echo "> 🐳 Start app service"
	@$(DOCKER_COMMAND) up zoity_app

docker-database:
	@echo "> 🐳 Start database service"
	@$(DOCKER_COMMAND) up zoity_postgres

# Docker commands

docker-up:
	@echo "> 🐳 Builds, (re)creates, starts, and attaches to containers for a service..."
	@$(DOCKER_COMMAND) up

docker-down:
	@echo "> 🐳 Stops containers and removes containers, networks, volumes, and images..."
	@$(DOCKER_COMMAND) down

docker-build:
	@echo "> 🐳 Build or rebuild services..."
	@$(DOCKER_COMMAND) build

docker-start:
	@echo "> 🐳 Start services..."
	@$(DOCKER_COMMAND) start

docker-stop:
	@echo "> 🐳 Stop services..."
	@$(DOCKER_COMMAND) stop

docker-restart:
	@echo "> 🐳 Restart service containers..."
	@$(DOCKER_COMMAND) restart

docker-logs:
	@echo "> 🐳 View output from containers..."
	@$(DOCKER_COMMAND) logs

docker-clean:
	@echo "> 🐳 Remove orphaned containers, volumes, and images...."
	@$(DOCKER_COMMAND) down -v
	@docker system prune -f
