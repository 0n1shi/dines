CC := $(shell which go)
BIN_DIR = bin
BIN = dines
SRC = cmd/dines/

.PHONY: build test

all: build

build:
	cd $(SRC) && $(CC) build

test:
	$(CC) test ./...
