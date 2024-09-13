# Makefile

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

.PHONY: help
all: help

.PHONY: dev
# dev:
# 	docker compose -f ./example/compose.yaml up --remove-orphans

.PHONY: test
test:
	go test -coverprofile=coverage.out -cover ./... && go tool cover -html=coverage.out -o coverage.html

.PHONY: docs
docs:
	go run $(GOBASE)/cmd/gomodeler docs

.PHONY: run
run:
	go run $(GOBASE)/cmd/gomodeler -e "teste1=true,teste2=false" -f ./examples/multiple-templates/envFile.yaml --template-path examples/multiple-templates/templates --log-level debug --output-path ./examples/multiple-templates/outputs

.PHONY: run-test
run-test:
	go run $(GOBASE)/cmd/gomodeler -f examples/complete/envFile.yaml --template-path examples/complete/templates --output-path examples/complete/outputs --log-level debug


.PHONY: build
build:
	go build -v -ldflags="-X 'main.Version=v.1.0-beta' -X 'main.commit=$(shell git rev-parse --short HEAD)' -X 'main.builtBy=$(shell id -u -n)' -X 'main.date=$(shell date)'" $(GOBASE)/cmd/gomodeler

.PHONY: version
version:
	go run -ldflags="-X 'main.version=v.1.0-beta' -X 'main.commit=$(shell git rev-parse --short HEAD)' -X 'main.builtBy=$(shell id -u -n)' -X 'main.date=$(shell date)'" $(GOBASE)/cmd/gomodeler version

.PHONY: help
help: Makefile
	@echo
	@echo "Usage: make [options]"
	@echo
	@echo "Options:"
	@echo "    build     Create binary file"
	@echo "    run       Run GoModeler"
	@echo "    version   Set version in go application"
	@echo "    Help	"
