.PHONY: build clean tool lint help

all: build

build:
	@go build -v .

tool:
	go vet ./...; true
	gofmt -w .

start:
	go run main.go

clean:
	rm -rf goRestServices
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make start: go run main.go"
	@echo "make clean: remove object files and cached files"