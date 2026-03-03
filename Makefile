DOCKER_IMAGE ?= hub.focot.cn/private/f-assets

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

.PHONY: build-frontend
build-frontend:
	cd frontend && pnpm i
	cd frontend && pnpm build

.PHONY: build-docker
build-docker: build-frontend
	docker build -t $(DOCKER_IMAGE):$(shell date +%Y%m%d)-$(shell git rev-parse HEAD | cut -c1-8) --push .

.PHONY: server-dev
server-dev:
	go run github.com/TiyaAnlite/F-Assests

.PHONY: web-dev
web-dev:
	cd frontend && pnpm dev