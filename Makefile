.DEFAULT_GOAL := all

SHELL := /bin/bash
.SHELLFLAGS := -euo pipefail -c

BINARY := envy
PLATFORMS := darwin linux
ARCHITECTURES := amd64

SOURCES := $(shell find . -type f -name '*.go')
OUT := target/
TARGET_PREFIX := $(OUT)$(BINARY)-
TARGETS := $(foreach platform,$(PLATFORMS),$(foreach architecture,$(ARCHITECTURES),$(TARGET_PREFIX)$(platform)-$(architecture)))
TARGET := $(TARGET_PREFIX)$(shell go env GOOS)-$(shell go env GOARCH)
CHECKSUMS := $(TARGETS:=.sha256)
SIGNATURES := $(TARGETS:=.asc)

GOOS_GOARCH = $(subst -, ,$(subst $(TARGET_PREFIX),,$@))
GOOS = $(firstword $(GOOS_GOARCH))
GOARCH = $(lastword $(GOOS_GOARCH))
LDFLAGS = -ldflags "-X main.Version=`cat $(VERSION)`"

VERSION := $(OUT)version
WRITE_VERSION := $(shell script/write-version $(VERSION))

$(VERSION):
	@script/write-version $(VERSION)

$(TARGETS): $(SOURCES) $(VERSION)
	@CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LDFLAGS) -o $@ cmd/envy/*.go

$(CHECKSUMS): %.sha256: %
	@(cd $(dir $<); shasum --algorithm 256 $(notdir $<)) > $@

$(SIGNATURES): %.asc: %
	@gpg --batch --armor --detach-sig --output $@ $<

all: get test

build: $(TARGET)

build-all: $(TARGETS)

check:
	@script/check

clean:
	@rm -Rf $(OUT)

get:
	@go get github.com/jstemmer/go-junit-report
	@dep ensure

install:
	@go install $(LDFLAGS)

release: $(CHECKSUMS) $(SIGNATURES)

test: build
	@go test -v ./test | tee >(go-junit-report >$(OUT)test-results.xml)

.PHONY: all build check clean get install release test
