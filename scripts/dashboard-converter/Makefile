FILES_TO_FMT=$(shell find . -name '*.go' -print)
TARGET=./build/dashboard-converter

all: build

build:
	@echo "Building $(TARGET) ..."
	@go build -o $(TARGET) ./...

test: build
	@echo "Testing..."
	@go test -v ./...

lint:
	@echo "Linting..."
	@golangci-lint run

vet:
	@echo "Vetting..."
	@go vet ./...

fmt:
	@gofmt -s -w $(FILES_TO_FMT)
	@goimports -w $(FILES_TO_FMT)

clean:
	@echo "Cleaning..."
	@go clean
	@rm -f $(TARGET)

.PHONY: all build test lint vet clean fmt
