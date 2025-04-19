BUILD_DIR=build
REGISTRY=ghcr.io/eyepipe/eye
GIT_COMMIT=$(shell git rev-parse --short HEAD)
BUILD_TIME=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
VERSION=0.1.1
LDFLAGS="-s -w -X github.com/eyepipe/eye/internal/pkg/buildinfo.BuildArgTime=$(BUILD_TIME) -X github.com/eyepipe/eye/internal/pkg/buildinfo.BuildArgGitCommit=$(GIT_COMMIT) -X github.com/eyepipe/eye/internal/pkg/buildinfo.BuildArgVersion=$(VERSION)"

help:
	@echo 'Available targets:'
	@echo '  make SIZE="1g" randfile'
	@echo ' '
	@echo '  make docker-build'
	@echo '  make docker-push'
	@echo ' '
	@echo '  make build-cli'
	@echo ' '
	@echo '  make db-up'
	@echo '  make db-down'
	@echo '  make NAME="create_pages" db-create'
	@echo ' '

clean:
	rm -rf ./$(BUILD_DIR)/*

build-cli: clean
build-cli:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/eye-cli-linux-amd64 -ldflags=$(LDFLAGS) cmd/cli/*.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/eye-cli-linux-arm64 -ldflags=$(LDFLAGS) cmd/cli/*.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/eye-cli-windows-amd64 -ldflags=$(LDFLAGS) cmd/cli/*.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/eye-cli-darwin-amd64 -ldflags=$(LDFLAGS) cmd/cli/*.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/eye-cli-darwin-arm64 -ldflags=$(LDFLAGS) cmd/cli/*.go
	@echo "============================="
	@echo "./build/eye-cli-darwin-arm64"

# make SIZE="10g" randfile
randfile: SIZE="1g"
randfile:
	mkfile -n ${SIZE} ${SIZE}.tmp

db-create: NAME=
db-create:
	goose -dir db/migrations create $(NAME) sql

db-up:
	go run cmd/server/*.go db up -dir db/migrations

db-down:
	go run cmd/server/*.go db down -dir db/migrations

docker-build:
	@echo "Build a Docker container for the current platform (locally)"
	$(MAKE) buildx ARG=--load

docker-push:
	@echo "Build Docker images for all available platforms and push them to the registry <$(REGISTRY)>"
	$(MAKE) buildx ARG="--push --platform linux/arm64,linux/amd64"

buildx: ARG=
buildx:
	@echo "building for the current platform..."
	@echo "docker buildx create --use"

	docker buildx build \
	$(ARG) \
	--no-cache \
	--build-arg LDFLAGS=$(LDFLAGS) \
	--build-arg GIT_COMMIT=$(GIT_COMMIT) \
	--build-arg VERSION=$(VERSION) \
	-t $(REGISTRY):$(VERSION) \
	-t $(REGISTRY):latest \
	-f docker/Dockerfile .

	@echo "👌 OK"
	@echo "docker run --rm -i ${REGISTRY}"
	@echo "docker run --rm -i --entrypoint server ${REGISTRY}"
