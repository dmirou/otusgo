PROJECT = github.com/dmirou/otusgo/hw07copyfile
HAS_LINT := $(shell command -v golangci-lint;)
GO_PACKAGES=$(shell go list $(PROJECT)/...)

all: vendor test lint

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

lint: bootstrap
	@echo "+ $@"
	@golangci-lint run ./...

.PHONY: all \
		vendor \
		bootstrap \
		lint
