APP_NAME := server
CMD_DIR := ./cmd/server
CONFIG_DIR := ./configs

.PHONY: help init proto wire build run test tidy clean release

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
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest

proto:
	buf generate

wire:
	wire $(CMD_DIR)

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
