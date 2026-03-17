.DEFAULT_GOAL := build

build-ui:
	cd ./cmd/web/ui && npm build
.PHONY: build-ui

build:
	go build
.PHONY: build