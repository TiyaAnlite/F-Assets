.PHONY: cli
cli: prepare build-cli

.PHONY: proto
proto:
	protoc --go_out=./pb pb/*.proto

.PHONY: prepare
prepare:
	mkdir -p build

.PHONY: build-cli
build-cli:
	go build -trimpath -v -o build/FAssetsCLI client/cli/*
