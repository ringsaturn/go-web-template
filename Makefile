.PHONY:build
build:
	cd cmd/app;make build

.PHONY:pb
pb:
	protoc -I=./api/ --go_out=./api/ --go-grpc_out=./api/ ./api/*.proto

tidy:
	rm go.sum
	go mod tidy

test:
	echo "test"
