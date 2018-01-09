TARGET = colander
CC = go build
GOARCH = amd64

# Symlink into GOPATH
BUILD_DIR=cmd/colander
CURRENT_DIR=$(shell pwd)

# Build the project
all: link linux

test:
	go test -race -v ./...

link:
	@BUILD_DIR=${BUILD_DIR}; \
	CURRENT_DIR=${CURRENT_DIR}; \

linux: 
	@cd ${BUILD_DIR}; \
	GOOS=linux GOARCH=${GOARCH} ${CC} -o ${TARGET}.o . ; \
	cd - >/dev/null

darwin:
	@cd ${BUILD_DIR}; \
	GOOS=darwin GOARCH=${GOARCH} ${CC} -o ${TARGET}.o . ; \
	cd - >/dev/null

clean:
	@cd ${BUILD_DIR}; \
	rm ${TARGET}.o ; \
	cd - >/dev/null

.PHONY: link linux
