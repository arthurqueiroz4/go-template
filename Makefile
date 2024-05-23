COVERAGE_FILE=coverage.out
COVERAGE_HTML=coverage.html
SRC_DIRS=./...
TEST_CMD=go test -coverprofile=$(COVERAGE_FILE) $(SRC_DIRS)
HTML_CMD=go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
SWAGGER_OUTPUT=docs

FILE ?= main.go
OUT_FILE ?= ./mock/mock_base.go

all: test coverage swagger

test:
	$(TEST_CMD)

coverage: test
	$(HTML_CMD)

swagger:
	swag init -g $(FILE) --output $(SWAGGER_OUTPUT)

mock:
	mockgen -source=$(FILE) -destination=$(OUT_FILE)

clean:
	rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)
	rm -rf $(SWAGGER_OUTPUT)

open-coverage: coverage
	$(HTML_CMD)
	@echo "Coverage report generated at $(COVERAGE_HTML)"
	@open $(COVERAGE_HTML)

run:
	@go run cmd/main.go

.PHONY: all test coverage clean open-coverage swagger mock
