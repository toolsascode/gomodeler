# Makefile

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

.PHONY: help
all: help

# .PHONY: dev
# dev:
# 	docker compose -f ./example/compose.yaml up --remove-orphans


.PHONY: dist-test
dist-test:
	goreleaser --snapshot --clean --skip=publish

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
	@echo
	@echo Cleaning up the outputs.
	rm -rf ./examples/complete/outputs ./examples/summary/outputs
	@echo Complete File testing
	go run $(GOBASE)/cmd/gomodeler -f examples/complete/envFile.yaml --template-path examples/complete/templates --output-path examples/complete/outputs --log-level debug
	@echo
	@echo Summary Files testing
	go run $(GOBASE)/cmd/gomodeler -f examples/summary/envFile.yaml --template-path ./.github/workflows/templates --output-path examples/summary/outputs --log-level debug


.PHONY: build
build:
	go build -v -ldflags="-X 'main.Version=v0.1.0-beta' -X 'main.commit=$(shell git rev-parse --short HEAD)' -X 'main.builtBy=$(shell id -u -n)' -X 'main.date=$(shell date)'" $(GOBASE)/cmd/gomodeler

.PHONY: version
version:
	go run -ldflags="-X 'main.version=v0.1.0-beta' -X 'main.commit=$(shell git rev-parse --short HEAD)' -X 'main.builtBy=$(shell id -u -n)' -X 'main.date=$(shell date)'" $(GOBASE)/cmd/gomodeler version
	@echo
	go run -ldflags="-X 'main.version=v0.1.0-beta' -X 'main.commit=$(shell git rev-parse --short HEAD)' -X 'main.builtBy=$(shell id -u -n)' -X 'main.date=$(shell date)'" $(GOBASE)/cmd/gomodeler -v
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
