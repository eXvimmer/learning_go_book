.DEFAULT_GOAL := build

.PHONY: build run run-client

build:
	go build

run: build
	./context_values

run-client:
	curl -b user=mustafa localhost:3000/?data=howdy
