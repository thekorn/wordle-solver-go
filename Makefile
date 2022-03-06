BINARY_NAME=main.out

GO := go1.18rc1
 
all: build test

build:
	$(GO) build -o ${BINARY_NAME} main.go
 
test:
	$(GO) test -coverprofile cover.out -v ./...

cover: test
	$(GO) tool cover -html=cover.out

run: run-naive

run-naive: build
	./${BINARY_NAME} -implementation naive -max 1
 
clean:
	go clean
	rm ${BINARY_NAME}