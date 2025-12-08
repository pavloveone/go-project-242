APP_NAME=hexlet-path-size
BIN=bin/${APP_NAME}

.PHONY: build run test

build:
	go build -o ${BIN} ./cmd/${APP_NAME}

run:
	./${BIN} $(ARGS)

test:
	go test -v ./...
