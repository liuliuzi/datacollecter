CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
TAG?=latest
COMMIT=`git rev-parse --short HEAD`

all: build

clean:
	@rm -rf bin/*

build:
	@cd cmd && go build  -tags netgo -o ../bin/datacollecter datacollecter.go

image: build
	@echo Building Shipyard image $(TAG)
	@docker build -t rcp/datacollecter:$(TAG) .

test: clean
	@godep go test -v ./...

.PHONY: all build clean  image test release