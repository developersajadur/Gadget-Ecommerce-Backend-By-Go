# Variables
APP_NAME=ecommerce
PORT?=8080

# Run the app (uses your run.sh script)
run:
	./scripts/run.sh

# Run the app without script (directly Go)
run-direct:
	@echo "🚀 Starting $(APP_NAME) on port $(PORT)"
	@PORT=$(PORT) go run ./cmd/main.go

# Run tests
test:
	@echo "🧪 Running tests..."
	go test ./... -cover

# Run database migrations
migrate:
	./scripts/migrate.sh

# Clean cache and binaries
clean:
	go clean
