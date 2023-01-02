.PHONY: all
all: worker

CMD_DIR := $(PWD)/cmd
PKG_DIR := $(PWD)/pkg

TOOLS_DIR := $(PWD)/tools
BUILD_DIR := $(PWD)/build

TOOLS_BIN_DIR := $(TOOLS_DIR)/bin

PATH := $(TOOLS_BIN_DIR):$(PATH)
export PATH

.PHONY: worker
worker:
	go build -mod=vendor -o $(BUILD_DIR)/$@ $(CMD_DIR)/$@/...

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

.PHONY: fmt
fmt:
	$(PWD)/scripts/gofmt.sh $(PWD)

.PHONY: lint
lint:
	golangci-lint run $(PWD)/...

.PHONY: vendor
vendor: vendor
	cd $(TOOLS_DIR) && $(MAKE) vendor
	go mod tidy && go mod vendor && go mod tidy && go mod verify

.PHONY: tools
tools:
	cd $(TOOLS_DIR) && $(MAKE)
