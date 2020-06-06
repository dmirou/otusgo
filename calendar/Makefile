PROJECT = github.com/dmirou/otusgo/calendar
APP = calendar
BIN = calendar

# Use the v0.0.0 tag for testing, it shouldn't clobber any release builds
RELEASE ?= v0.11
PROTOC_IMAGE ?= dmirou/protoc:0.1.1

PWD = $(shell pwd)
PROTOC_RUN = docker run --rm -v $(PWD):/contracts -w /contracts \
	-u $(shell id -u):$(shell id -g) $(PROTOC_IMAGE) -I./pkg/contracts/

REPO_INFO=$(shell git config --get remote.origin.url)
REPO_BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
RELEASE_DATE=$(shell date +%FT%T%Z)

ifndef REPO_COMMIT
REPO_COMMIT = git-$(shell git rev-parse --short HEAD)
endif

HAS_LINT := $(shell command -v golangci-lint;)

LDFLAGS = "-s -w \
	-X $(PROJECT)/pkg/version.RELEASE=$(RELEASE) \
	-X $(PROJECT)/pkg/version.DATE=$(RELEASE_DATE) \
	-X $(PROJECT)/pkg/version.REPO=$(REPO_INFO) \
	-X $(PROJECT)/pkg/version.COMMIT=$(REPO_COMMIT) \
	-X $(PROJECT)/pkg/version.BRANCH=$(REPO_BRANCH)"

GO_PACKAGES=$(shell go list $(PROJECT)/...)

CORE_SERVER_ADDR?="127.0.0.1:9000"

all: build

bootstrap:
ifndef HAS_LINT
	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint@v1.20.0
endif

vendor: bootstrap
	@echo "+ $@"
	@go mod tidy

fmt:
	@echo "+ $@"
	@gofmt -w ./pkg/

test:
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"go test -race -cover {{.Dir}}"{{end}}' $(GO_PACKAGES) | xargs -L 1 sh -c

lint: bootstrap fmt
	@echo "+ $@"
	@golangci-lint run ./...

contracts:
	@echo "+ $@"
	@$(PROTOC_RUN) --go_out=tag=true,plugins=grpc,paths=source_relative:./pkg/contracts/ \
		./pkg/contracts/request/request.proto
	@$(PROTOC_RUN) --go_out=tag=true,plugins=grpc,paths=source_relative:./pkg/contracts/ \
		./pkg/contracts/event/event.proto

build: contracts vendor lint test
	@echo "+ $@"
	@go build -a -ldflags $(LDFLAGS) -o bin/$(BIN) $(PROJECT)/cmd/calendar

run: build
	@echo "+ $@"
	@bin/$(BIN) --config ./config/default.config

run-bg: build
	@echo "+ $@"
	@bin/$(BIN) --config ./config/default.config > /dev/null 2>&1 &

sleep:
	@echo "+ $@"
	@sleep 3

integration-test: stop-service run-bg sleep integration-test-only stop-service
	@echo "+ $@"

integration-test-only:
	@echo "+ $@"
	@CORE_SERVER_ADDR="$(CORE_SERVER_ADDR)" go test -tags=integration -v -race -cover -bench=. -benchmem ./integration/...

stop-service:
	@echo "+ $@"
	@killall -9 $(BIN)
	@sleep 1

.PHONY: all \
		bootstrap \
		build \
		contracts \
		lint \
		run \
		vendor
