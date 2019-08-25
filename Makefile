TAG=$(shell git describe --tags --always)
VERSION=$(TAG:v%=%)
NAME=pingpong
REPO=shawntoffel/$(NAME)
GO=GO111MODULE=on go
BUILD=GOARCH=amd64 $(GO) build -ldflags="-s -w -X main.Version=$(VERSION)" 
PROTOFILES=$(wildcard api/protobuf-spec/*/*.proto)
PBFILES=$(patsubst %.proto,%.pb.go, $(PROTOFILES))

.PHONY: all deps test build clean $(PROTOFILES)

all: deps test build 
deps:
	$(GO) mod download

test:
	$(GO) vet ./...
	$(GO) test -v -race ./...

build:
	$(BUILD) -o bin/$(NAME)-$(VERSION) ./cmd/...

clean:
	@find bin -type f ! -name '*.toml' -delete -print
