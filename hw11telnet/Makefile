PROJECT = github.com/dmirou/otusgo/hw11telnet
BIN = gotelnet

HAS_LINT := $(shell command -v golangci-lint;)
GO_PACKAGES=$(shell go list $(PROJECT)/...)

GOOS ?= linux
GOARCH ?= amd64

all: run

run: build
	@echo "+ $@"
	bin/$(GOOS)-$(GOARCH)/$(BIN)

build: clean lint compile

clean:
	@echo "+ $@"
	@rm -f bin/$(GOOS)-$(GOARCH)/$(BIN)

lint: bootstrap
	@echo "+ $@"
	@golangci-lint run ./...

bootstrap:
ifndef HAS_LINT
	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
endif

compile: vendor
	@echo "+ $@"
	@CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -a -installsuffix cgo \
			-o bin/$(GOOS)-$(GOARCH)/$(BIN) $(PROJECT)/cmd

vendor: bootstrap
	@echo "+ $@"
	@go mod tidy

demo: clean lint compile
	@echo "+ $@"
	bin/$(GOOS)-$(GOARCH)/$(BIN) google.com 80

.PHONY: all \
		clean \
		compile \
		demo \
		run \
		vendor \
		bootstrap \
		lint
