.PHONY: build
build:
	go build -o Builds-List.exe -ldflags="-s -w" -v ./cmd/buildlist 

.DEFAULT_GOAL := build