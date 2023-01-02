.PHONY: all
all: \
	install_gci \
	install_gofumpt \
	install_goimports \
	install_golangci-lint

BIN_DIR := $(PWD)/bin

.PHONY: vendor
vendor: vendor
	go mod tidy && go mod vendor && go mod tidy && go mod verify

.PHONY: install_gci
install_gci:
	GOBIN=$(BIN_DIR) go install -ldflags '-s -w' ./vendor/github.com/daixiang0/gci

.PHONY: install_gofumpt
install_gofumpt:
	GOBIN=$(BIN_DIR) go install -ldflags '-s -w' ./vendor/mvdan.cc/gofumpt

.PHONY: install_goimports
install_goimports:
	GOBIN=$(BIN_DIR) go install -ldflags '-s -w' ./vendor/golang.org/x/tools/cmd/goimports

.PHONY: install_golangci-lint
install_golangci-lint:
	GOBIN=$(BIN_DIR) go install -ldflags '-s -w' ./vendor/github.com/golangci/golangci-lint/cmd/golangci-lint
