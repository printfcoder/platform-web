
GOPATH:=$(shell go env GOPATH)

.PHONY: build
build: proto

	go build -o platform-web main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user-srv:latest
