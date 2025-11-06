SERVER_NAME := "givemegoodcoffee server"
SERVER_BINARY := ./bin/server
SERVER_ENTRY_POINT := ./cmd/server/
SERVER_PID_FILE := .app.pid

build:
	@echo "Building $(SERVER_NAME)"
	@mkdir -p bin
	@go build -o $(SERVER_BINARY) $(SERVER_ENTRY_POINT)

run: build
	@echo "Starting $(SERVER_NAME)..."
	@$(SERVER_BINARY)

start: build
	@echo "Starting $(SERVER_NAME) in background"
	@$(SERVER_BINARY) & echo $$! > $(SERVER_PID_FILE)
	@echo "$(SERVER_NAME) started with PID $$(cat $(SERVER_PID_FILE))"

stop:
	@if [ -f $(SERVER_PID_FILE) ]; then \
		echo "Stopping $(SERVER_NAME) with PID $$(cat $(SERVER_PID_FILE))"; \
		kill $$(cat $(SERVER_PID_FILE)) 2>/dev/null || true; \
		rm -f $(SERVER_PID_FILE); \
		echo "$(SERVER_NAME) stopped"; \
	else \
		echo "No PID file found. $(SERVER_NAME) is probably not running."; \
	fi

restart: stop start
	@echo "$(SERVER_NAME) restarted successfully"

clean:
	@echo "Cleaning $(SERVER_NAME) binaries"
	@rm -rf bin/
