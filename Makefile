# Go related variables.
PROJECTNAME=$(shell basename "$(PWD)")
GOBASE=$(shell pwd)
GOFILES=$(wildcard *.go)
BIN=$(GOBASE)/.bin
export GOBIN=$(BIN)

ifneq (,$(wildcard ./.env))
	include .env
	export
endif

# Tools

GO_AIR = $(BIN)/air

$(BIN)/air: PACKAGE=github.com/cosmtrek/air@latest

$(BIN):
	@mkdir -p $@

$(BIN)/%: | $(BIN)
	go install $(PACKAGE)

dev: hooks $(GO_AIR)
	$(GO_AIR)

hooks: build 
	@chmod +x .githooks/pre-commit
	@git config core.hooksPath .githooks
	@echo "Git hooks are installed"

build:
	@echo "Building..."
	@GOOS=linux go build -ldflags="-s -w" -o bin/app ./cmd/app/
	@GOOS=linux go build -ldflags="-s -w" -o bin/crawler ./cmd/crawler/
	@echo "Build done"

start: build
	@echo "Starting..."
	@./dist/app
