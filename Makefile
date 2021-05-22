.PHONY:build
build:
	cd cmd/app;make build

.PHONY:pb
pb:
	protoc -I=./api/ --go_out=./api/ --go-grpc_out=./api/ --doc_out=./api/ --doc_opt=markdown,docs.md ./api/*.proto

tidy:
	rm go.sum
	go mod tidy

test:
	echo "test"

install:
	go mod download
	go get github.com/google/wire/cmd/wire
	go get github.com/golang/protobuf/protoc-gen-go
	go get github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
