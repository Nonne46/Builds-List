.PHONY: build
build:
	go build -ldflags="-s -w" -v ./cmd/buildlist 

.DEFAULT_GOAL := build