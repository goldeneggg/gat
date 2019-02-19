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
	@go tool vet --all -shadow ./*.go
	@go tool vet -all -shadow ./client
	@go tool vet -all -shadow ./client/http

dep-save:
	@echo "Run godep save"
	@save -v $(SAVE_TARGET)

dep-saved-build: dep-save build

lint:
	@echo "Linting"
	@${GOBIN}/golint $(PGM_PATH)
	@${GOBIN}/golint $(PGM_PATH)/client
	@${GOBIN}/golint $(PGM_PATH)/client/http

release:
	@echo "Releasing"
	@./scripts/release.sh

publish: release
	@echo "Publishing releases to github"
	@./scripts/publish.sh

formula:
	@echo "Generating formula"
	@./scripts/publish.sh formula-only
