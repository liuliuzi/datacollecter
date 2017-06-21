CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
TAG?=latest
COMMIT=`git rev-parse --short HEAD`

all: build

clean:
	@rm -rf controller/controller

build:
	@cd controller && godep go build -a -tags "netgo static_build" -installsuffix netgo -ldflags "-w -X github.com/shipyard/shipyard/version.GitCommit=$(COMMIT)" .

image: build
	@echo Building Shipyard image $(TAG)
	@cd controller && docker build -t shipyard/shipyard:$(TAG) .

release: build image
	@docker push shipyard/shipyard:$(TAG)

test: clean
	@godep go test -v ./...

.PHONY: all build clean  image test release