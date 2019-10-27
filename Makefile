#SHELL = /bin/bash
# GO parameters
GOCMD=go
BUILDENV=GOTRACEBACK=none CGO_ENABLED=0
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

BINARY_NAME=golden-headed-quetzal

LDFLAGS=-ldflags "-w -s"

build:
	@$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) *.go

build-debug:
	@$(GOBUILD) -o $(BINARY_NAME) *.go

clean:
	@rm -rf golden-headed-quetzal

run: build
	./golden-headed-quetzal

run-debug: build-debug
	./golden-headed-quetzal

fmt:
	@gofmt -w *.go
	@gofmt -w **/**.go

deps:
	@go get -v ./...

tests:
	@go test $(go list ./... | grep -v /vendor/)