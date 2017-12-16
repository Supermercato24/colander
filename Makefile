BINARY = archive
CC = go build
GOARCH = amd64

# Symlink into GOPATH
BUILD_DIR=cmd/centrifuge
CURRENT_DIR=$(shell pwd)

# Build the project
all: link linux

link:
	@BUILD_DIR=${BUILD_DIR}; \
	CURRENT_DIR=${CURRENT_DIR}; \

linux: 
	@cd ${BUILD_DIR}; \
	GOOS=linux GOARCH=${GOARCH} ${CC} -o ${BINARY}.o . ; \
	cd - >/dev/null

darwin:
	@cd ${BUILD_DIR}; \
	GOOS=darwin GOARCH=${GOARCH} ${CC} -o ${BINARY}.o . ; \
	cd - >/dev/null

clean:
	@cd ${BUILD_DIR}; \
	rm ${BINARY}.o ; \
	cd - >/dev/null

.PHONY: link linux
