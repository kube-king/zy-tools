VERSION=v1.0.0
# OS
GOOS=linux
# go proxy
GOPROXY=https://goproxy.cn,direct
# go package module
GO111MODULE=on
#Multi Arch
TARGET_ARCHS ?= amd64 arm64
# Docker Image Tag
IMAGE_TAG=sit-registry.faw.cn
# APP
APP=zy-tools

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...


.PHONY: lint
lint: ## Run golangci-lint.
	golangci-lint run

.PHONY: test
test: ## Run tests.
	go test ./... -coverprofile=coverage.out

.PHONY: run
run: fmt vet ## Run the binary.
	go run ./cmd/zy_tools/main.go


.PHONY: build
build: ## Build the binary.
	go build -o ./build/zy_tools/zy_tools ./cmd/zy_tools/main.go

.PHONY: docker-build
docker-build: ## Build the docker image.
	docker build -t ${IMAGE_TAG}/qkp-system/faw/${APP}:${VERSION} -f build/zy_tools/Dockerfile .

.PHONY: docker-push
docker-push: ## Push the docker image.
	docker push ${IMAGE_TAG}/qkp-system/faw/${APP}:${VERSION}

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $()]]]]]]


.PHONY: docker-buildx
docker-buildx: ## Build the docker image.
	docker buildx build --push -f build/zy_tools/Dockerfile  $(foreach arch,${TARGET_ARCHS},--platform=linux/${arch}) $(foreach tag,${IMAGE_TAG},--tag=${tag}/qkp-system/faw/${APP}:${VERSION})  .

.PHONY: docker-create-buildx
docker-create-buildx: ## Create Builder Context.
	docker buildx create  --name=${APP} --driver docker-container --driver-opt image=moby/buildkit:master

.PHONY: docker-delete-buildx
docker-delete-buildx: ## Delete Builder Context.
	docker buildx rm ${APP}


# Builder middleware-api Multi Arch
#docker-buildx: fmt vet
#	docker buildx build --builder=${APP}  --push -f build/zy_tools/Dockerfile  $(foreach arch,${TARGET_ARCHS},--platform=linux/${arch}) $(foreach tag,${IMAGE_TAG},--tag=${tag}/qkp-system/faw/${APP}:${VERSION})  .
#
## Create Builder Context
#docker-create-buildx:
#	docker buildx create  --name=${APP} --driver docker-container --driver-opt image=moby/buildkit:master

# start dev
.PHONY: dev
dev:
	air -c air.conf


