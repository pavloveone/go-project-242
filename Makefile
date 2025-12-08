APP_NAME=hexlet-path-size
BIN=bin/${APP_NAME}

.PHONY: build 


build:
	go build -o ${BIN} ./cmd/${APP_NAME}
	./${BIN}

