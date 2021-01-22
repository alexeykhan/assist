export GO111MODULE=on
export GOSUMDB=off

BIN := $(CURDIR)/bin
CMD := $(CURDIR)/cmd
CLI_BIN_NAME := oracle
CLI_CMD_NAME := console
LINTER_TAG := v1.32.2
LINTER_BIN := $(BIN)/golangci-lint
CHECKLICENSE := $(CURDIR)/checklicense.sh
LOGS_DIR := $(CURDIR)/logs
TESTS_LOGS := $(LOGS_DIR)/tests.log
LINTER_LOGS := $(LOGS_DIR)/linter.log

.PHONY: all
all: lint test

.PHONY: lint
lint:
	@echo "# Checking license headers ..."
	@bash $(CHECKLICENSE) | tee $(LINTER_LOGS)
	@echo "# Checking code with linters ..."
	@$(LINTER_BIN) run --config=.golangci.yaml ./... | tee $(LINTER_LOGS)
	@[ ! -s $(LINTER_LOGS) ]

.PHONY: test
test:
	@echo "# Running tests ..."
	@go test -race ./... | tee $(TESTS_LOGS)

.PHONY: cover
cover:
	@echo "# Running coverage tests ..."
	@go test -race -covermode=atomic -coverprofile=cover.out -coverpkg=./... ./...
	@echo "# See coverage results in ./cover.html file"
	@go tool cover -html=cover.out -o cover.html

.PHONY: generate
generate:
	@echo "# Generating stuff ..."
	@go generate ./...

.PHONY: linter
linter:
	@echo "# Installing golangci-lint ..."
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin $(LINTER_TAG)

.PHONY: build-cli
build-cli:
	@echo "# Building cli ..."
	@go build -o $(BIN)/$(CLI_BIN_NAME) $(CMD)/$(CLI_CMD_NAME)
