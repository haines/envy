.DEFAULT_GOAL := build

BINARY := envy
SOURCES := $(shell find . -type f -name '*.go')
OUT := target/
TARGET := $(OUT)$(BINARY)
VERSION := $(OUT)version

DOCKER_REPO := ahaines/$(BINARY)

LDFLAGS = -ldflags "-X github.com/haines/envy/cmd.Version=`cat $(VERSION)`"

WRITE_VERSION := $(shell ./write-version $(VERSION))

$(TARGET): $(SOURCES) $(VERSION)
	@go build $(LDFLAGS) -o $(TARGET)

all: get build

build: $(TARGET)

clean:
	@rm -Rf $(OUT)

get:
	@go get ./...

install:
	@go install $(LDFLAGS)

.PHONY: all build clean get install
