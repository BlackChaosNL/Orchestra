.DEFAULT_GOAL := build-linux

.PHONY: remove-node-modules
remove-node-modules:
	@rm -r ./cmd/web/ui/node_modules || true

.PHONY: remove-dist
remove-dist:
	@rm -r ./cmd/web/ui/dist || true

.PHONY: build-ui
build-ui:
	@cd ./cmd/web/ui && \
	npm i && \
	npm run build

.PHONY: build-linux
build-linux:
	make remove-node-modules
	make remove-dist
	make build-ui
	make remove-node-modules

	@GOOS=linux go build -ldflags="-s -w" -o ./Orchestra ./*.go
