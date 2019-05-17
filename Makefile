NAME := go-git-fame

VERSION := v1.0.0

LDFLAGS  := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\""

GO_SRC_DIRS := $(shell \
	find . -name "*.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)

GO_TEST_DIRS := $(shell \
	find . -name "*_test.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)

NOVENDOR := $(shell go list ./... | grep -v vendor)

TEST_PATTERN?=.

TEST_OPTIONS?=-race -covermode=atomic -coverprofile=coverage.txt

.DEFAULT_GOAL := bin/$(NAME)

bin/$(NAME): $(GO_SRC_DIRS)
	go build $(LDFLAGS) -o bin/$(NAME)

.PHONY: setup
setup:  ## Installs all of the build and lint dependencies
	go get -u gopkg.in/alecthomas/gometalinter.v2
	go get -u github.com/golang/dep/cmd/dep
	go get -u golang.org/x/tools/cmd/cover
	go get -u golang.org/x/tools/cmd/goimports
	gometalinter --install --update

.PHONY: install
install:  ## runs install
	go install $(LDFLAGS)

.PHONY: test
test: ## runs test with coverage (doesn't produce a report)
	echo 'mode: atomic' > bin/coverage.txt && go list ./... | xargs -n1 -I{} sh -c 'go test -covermode=atomic -coverprofile=bin/coverage.tmp {} && tail -n +2 bin/coverage.tmp >> bin/coverage.txt' && rm bin/coverage.tmp

.PHONY: cover
cover: test ## Run all the tests and opens the coverage report
		go tool cover -html=bin/coverage.txt

.PHONY: fmt
fmt: ## runs fmt
	go fmt $(NOVENDOR)
.PHONY: clean
clean: ## clean up
	rm -rf bin/*
	rm -rf vendor/*

.PHONY: help
help:  ## displays this message
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:  ## displays the version
	@echo $(VERSION)
