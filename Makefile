#GO ?= go
#GODEP ?= godep
#GOLINT ?= golint
NAME := gat
SRCS := $(shell find . -type f -name '*.go' | \grep -v 'vendor')
PGM_PATH := 'github.com/goldeneggg/gat'
SAVE_TARGET := ./...
PROFDIR := ./.profile
PROFTARGET := ./client

.DEFAULT_GOAL := bin/$(NAME)

all: build

mod-dl:
	@GO111MODULE=on go mod download

bin/$(NAME): $(SRCS)
	@echo "Building bin/$(NAME)"
	@go build -o bin/$(NAME) $(PGM_PATH)

.PHONY: test
test:
	@echo "Testing"
	@go test -race -v $(PGM_PATH)
	@go test -race -v $(PGM_PATH)/client...

.PHONY: prof
prof:
	@[ ! -d $(PROFDIR) ] && mkdir $(PROFDIR); go test -bench . -benchmem -blockprofile $(PROFDIR)/block.out -cover -coverprofile $(PROFDIR)/cover.out -cpuprofile $(PROFDIR)/cpu.out -memprofile $(PROFDIR)/mem.out $(PROFTARGET)

.PHONY: vet
vet:
	@echo "Vetting"
	@go vet -all -shadow ./*.go
	@go vet -all -shadow ./client
	@go vet -all -shadow ./client/http

.PHONY: lint
lint:
	@echo "Linting"
	@golint -set_exit_status $(PGM_PATH)
	@golint -set_exit_status $(PGM_PATH)/client
	@golint -set_exit_status $(PGM_PATH)/client/http

.PHONY: validate
validate: vet lint

lint-travis:
	@travis lint .travis.yml

test-goreleaser:
	@echo "Testing goreleaser"
	@goreleaser release --snapshot --skip-publish --rm-dist

test-goreleaser-on-ci:
	@echo "Testing goreleaser (on CI)"
	@./goreleaser release --snapshot --skip-publish --rm-dist

.PHONY: ci
ci: test lint test-goreleaser-on-ci

.PHONY: release
release:
	@echo "Releasing"
	@./scripts/release.sh

.PHONY: publish
publish: release
	@echo "Publishing releases to github"
	@./scripts/publish.sh

.PHONY: formula
formula:
	@echo "Generating formula"
	@./scripts/publish.sh formula-only
