PROJECT = github.com/dmirou/otusgo/hw08envdir
BIN = goenv

HAS_LINT := $(shell command -v golangci-lint;)
GO_PACKAGES=$(shell go list $(PROJECT)/...)

GOOS ?= linux
GOARCH ?= amd64

all: run

clean:
	@echo "+ $@"
	@rm -f bin/$(GOOS)-$(GOARCH)/$(BIN)

run: clean lint test compile
	@echo "+ $@"
	bin/$(GOOS)-$(GOARCH)/$(BIN)

compile: vendor
	@echo "+ $@"
	@CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -a -installsuffix cgo \
			-o bin/$(GOOS)-$(GOARCH)/$(BIN) $(PROJECT)/cmd

vendor: bootstrap
	@echo "+ $@"
	@go mod tidy

bootstrap:
ifndef HAS_LINT
	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
endif

test:
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"go test -race -cover {{.Dir}}"{{end}}' $(GO_PACKAGES) | xargs -L 1 sh -c

demo: clean lint test compile
	@echo "+ $@"
	bin/$(GOOS)-$(GOARCH)/$(BIN) ./pkg/os/testenvdir/ printenv

lint: bootstrap
	@echo "+ $@"
	@golangci-lint run ./...

.PHONY: all \
		clean \
		compile \
		demo \
		run \
		vendor \
		bootstrap \
		lint
