.DEFAULT_GOAL := build-linux

.PHONY: clean-ui
clean-ui:
	@rm -r ./cmd/web/ui/node_modules
	@rm -r ./cmd/web/ui/dist

.PHONY: build-ui
build-ui:
	@cd ./cmd/web/ui && \
	npm i && \
	npm run build

.PHONY: build-linux
build-linux:
	@GOOS=linux go build -ldflags="-s -w" -o ./Orchestra ./*.go
