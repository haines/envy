.DEFAULT_GOAL := all

SHELL := /bin/bash
.SHELLFLAGS := -euo pipefail -c

BINARY := envy
SOURCES := $(shell find . -type f -name '*.go')
OUT := target/
TARGET := $(OUT)$(BINARY)
VERSION := $(OUT)version

DOCKER_REPO := ahaines/$(BINARY)

LDFLAGS = -ldflags "-X main.Version=`cat $(VERSION)`"

WRITE_VERSION := $(shell script/write-version $(VERSION))

$(VERSION):
	@script/write-version $(VERSION)

$(TARGET): $(SOURCES) $(VERSION)
	@go build $(LDFLAGS) -o $(TARGET) cmd/envy/*.go

all: get build test

build: $(TARGET)

check:
	@script/check

clean:
	@rm -Rf $(OUT)

get:
	@go get github.com/jstemmer/go-junit-report
	@dep ensure

install:
	@go install $(LDFLAGS)

test: build
	@go test -v ./test | tee >(go-junit-report >$(OUT)/test-results.xml)

.PHONY: all build check clean get install test
