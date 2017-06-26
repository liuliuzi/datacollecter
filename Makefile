CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
TAG?=latest
COMMIT=`git rev-parse --short HEAD`
GOLANG_VERSION?=1.7
VERSION?=v2.1.0
GIT_COMMIT:=$(shell git rev-parse --short HEAD)
APP_LDFLAGS=-w -X github.com/liuliuzi/datacollecter/version.AppVersion=$(VERSION) -X github.com/liuliuzi/datacollecter/version.GitCommit=$(GIT_COMMIT)

all: build

clean:
	@rm -rf bin/*

build:
	@cd cmd && go build  -tags netgo -ldflags "$(APP_LDFLAGS)" -o ../bin/datacollecter datacollecter.go

image: build
	@echo Building Shipyard image $(TAG)
	@docker build -t rcp/datacollecter:$(TAG) .

test: clean
	@godep go test -v ./...

.PHONY: all build clean  image test