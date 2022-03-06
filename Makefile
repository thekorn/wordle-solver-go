BINARY_NAME=main.out

GO := go1.18rc1
 
all: build test

build:
	$(GO) build -o ${BINARY_NAME} main.go
 
test:
	$(GO) test -coverprofile cover.out -v ./...

run: build outdir
	./${BINARY_NAME}
 
clean:
	go clean
	rm ${BINARY_NAME}