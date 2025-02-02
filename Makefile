BINARY_NAME=nyx
BUILD_DIR=bin
ENTRYPOINT=main.go

.PHONY: all
all: build

.PHONY: build
build:
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(ENTRYPOINT)

.PHONY: run
run: build
	@$(BUILD_DIR)/$(BINARY_NAME) <args>

.PHONY: test
test:
	@go test ./...

.PHONY: clean
clean:
	@rm -rf $(BUILD_DIR)

.PHONY: install
install: build
	@go install ./...

