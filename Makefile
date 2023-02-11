BINDIR      := $(CURDIR)/bin
DIST_DIRS   := find * -type d -exec
TARGETS     := linux/amd64 windows/amd64
TARGET_OBJS ?= linux-amd64.tar.gz linux-amd64.tar.gz.sha256 linux-amd64.tar.gz.sha256sum windows-amd64.zip windows-amd64.zip.sha256 windows-amd64.zip.sha256sum
BINNAME     ?= alerter

GOPATH        = $(shell go env GOPATH)
DEP           = $(GOPATH)/bin/dep
GOX           = $(GOPATH)/bin/gox
GOIMPORTS     = $(GOPATH)/bin/goimports
ARCH          = $(shell uname -p)

# go option
PKG        := ./...
TAGS       :=
TESTS      := .
TESTFLAGS  :=
LDFLAGS    := -w -s
GOFLAGS    :=
SRC        := $(shell find . -type f -name '*.go' -print)

# Required for globs to work correctly
SHELL      = /bin/bash

GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_SHA    = $(shell git rev-parse --short HEAD)
GIT_TAG    = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
GIT_DIRTY  = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")

ifdef VERSION
	BINARY_VERSION = $(VERSION)
endif
BINARY_VERSION ?= ${GIT_TAG}

# Only set Version if building a tag or VERSION is set
ifneq ($(BINARY_VERSION),)
	LDFLAGS += -X github.com/wooos/alerter/internal/pkg/version.version=${BINARY_VERSION}
endif

LDFLAGS += -X github.com/wooos/alerter/internal/pkg/version.metadata=${VERSION_METADATA}
LDFLAGS += -X github.com/wooos/alerter/internal/pkg/version.gitCommit=${GIT_COMMIT}
LDFLAGS += -X github.com/wooos/alerter/internal/pkg/version.gitTreeState=${GIT_DIRTY}
LDFLAGS += $(EXT_LDFLAGS)

.PHONY: all
all: build build-docker-image

# ------------------------------------------------------------------------------
#  build

.PHONY: build
build: $(BINDIR)/$(BINNAME)

$(BINDIR)/$(BINNAME): $(SRC)
	GO111MODULE=on go build $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o $(BINDIR)/$(BINNAME) ./cmd/alerter


.PHONY: build-docker-image
build-docker-image:
	docker build -t wooos/alerter:$(BINARY_VERSION) -f Dockerfile .

# ------------------------------------------------------------------------------
#  test

.PHONY: test
test: test-unit

.PHONY: test-unit
test-unit:
	@echo
	@echo "==> Running unit tests <=="
	GO111MODULE=on go test $(GOFLAGS) -run $(TESTS) $(PKG) $(TESTFLAGS)

.PHONY: format
format: $(GOIMPORTS)
	@echo
	@echo "==> Running format <=="
	GO111MODULE=on go list -f '{{.Dir}}' ./... | xargs $(GOIMPORTS) -w -local ./

# ------------------------------------------------------------------------------
#  #  dependencies

# If go get is run from inside the project directory it will add the dependencies
# # to the go.mod file. To avoid that we change to a directory without a go.mod file
# # when downloading the following dependencies
#
$(GOX):
	(cd /; GO111MODULE=on go get -u github.com/mitchellh/gox)

$(GOIMPORTS):
	(cd /; GO111MODULE=on go get -u golang.org/x/tools/cmd/goimports)

.PHONY: clean
clean:
	@rm -rf $(BINDIR) ./_dist

.PHONY: info
info:
	@echo "Version:           ${VERSION}"
	@echo "Git Tag:           ${GIT_TAG}"
	@echo "Git Commit:        ${GIT_COMMIT}"
	@echo "Git Tree State:    ${GIT_DIRTY}"