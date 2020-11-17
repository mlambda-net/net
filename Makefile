GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=monads
LINTER=golangci-lint

all: test build

generate:
	protoc -I=. -I=${GOPATH}/src --gogoslick_out=./pkg/core --go-grpc_out=./pkg/core  net.proto

test:
	$(GOTEST) ./... -v

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

lint:
	$(LINTER) run
