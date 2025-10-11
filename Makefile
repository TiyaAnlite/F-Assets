.PHONY: cli
cli: prepare build-cli

.PHONY: prepare
prepare:
	mkdir -p build

.PHONY: build-cli
build-cli:
	go build -C client/cli -trimpath -v -o ../../build/FAssetsCLI
