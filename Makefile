all: build test lint

build:
	go build ./... && goapp build ./...

test:
	go test ./... && goapp test ./...

lint:
	-golint ./... | grep -vE "^(vendor|logutil|set)/"
