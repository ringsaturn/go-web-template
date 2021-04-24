VERSION=$(shell git describe --tags --always)
BRANCH=$(shell git symbolic-ref -q --short HEAD)
REVISION=$(shell git rev-parse --short HEAD)
BUILD_DATE=$(shell date +%FT%T%z)

.PHONY:build
build:
	wire
	go build -ldflags "-X main.Version=$(VERSION) -X main.Branch=$(BRANCH) -X main.Revision=$(REVISION) -X main.BuildDate=$(BUILD_DATE)" -o ./../../build
