# Makefile

.PHONY: test run latest

LOG_DIR := logs
NOW := $(shell date +%Y-%m-%d-%H-%M-%S)
TEST_lOG := $(LOG_DIR)/$(NOW)-test.log

test:
	@mkdir -p $(LOG_DIR)
	@echo "Running tests and saving to $(TEST_lOG)"
	@go test ./tests/... 2>&1 | tee $(TEST_lOG)

run:
	@go run latest/*.go

latest:
	@echo "Symlink points to: $(shell readlink latest)"

