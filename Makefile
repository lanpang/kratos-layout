# Default to the only command under ./cmd so cloned/new repos do not keep
# a stale ./cmd/server path. Multi-command repos can override APP_NAME/CMD_DIR.
CMD_DIRS := $(sort $(wildcard ./cmd/*))
APP_NAME ?= $(if $(filter 1,$(words $(CMD_DIRS))),$(notdir $(firstword $(CMD_DIRS))),server)
CMD_DIR ?= ./cmd/$(APP_NAME)
CONFIG_DIR := ./configs
BUF ?= buf
WIRE ?= go tool wire

.PHONY: help init proto wire check-cmd build run test tidy clean release

help:
	@echo "Available targets:"
	@echo "  init     Install project codegen tools"
	@echo "  proto    Generate protobuf, gRPC, HTTP, validation, and docs"
	@echo "  wire     Generate compile-time dependency injection code"
	@echo "  build    Generate code and build $(APP_NAME)"
	@echo "  run      Run $(APP_NAME) with $(CONFIG_DIR)"
	@echo "  test     Run Go tests"
	@echo "  tidy     Tidy Go modules"
	@echo "  clean    Remove local build output"
	@echo "  release  Build release snapshot with GoReleaser"

init:
	go install github.com/bufbuild/buf/cmd/buf@latest
	go get -tool github.com/google/wire/cmd/wire@v0.7.0
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest

proto:
	$(BUF) generate
	cd internal/conf && $(BUF) generate

check-cmd:
	@test -d "$(CMD_DIR)" || { \
		echo "CMD_DIR '$(CMD_DIR)' does not exist."; \
		echo "Available command dirs: $(CMD_DIRS)"; \
		echo "Set APP_NAME=<name> or CMD_DIR=./cmd/<name>."; \
		exit 1; \
	}

wire: check-cmd
	$(WIRE) $(CMD_DIR)

build: proto wire
	go build -o bin/$(APP_NAME) $(CMD_DIR)

run: wire
	go run $(CMD_DIR) -conf $(CONFIG_DIR)

test:
	go test ./...

tidy:
	go mod tidy

clean:
	rm -rf bin dist

release:
	goreleaser build --snapshot --clean
