#GO ?= go
#GODEP ?= godep
#GOLINT ?= golint
BINNAME := gat
PGM_PATH := 'github.com/goldeneggg/gat'
SAVE_TARGET := ./...
PROFDIR := ./.profile
PROFTARGET := ./client

all: build

build:
	@echo "Building ${GOBIN}/$(BINNAME)"
	@GO15VENDOREXPERIMENT=1 godep go build -o ${GOBIN}/$(BINNAME) $(PGM_PATH)

test-all:
	@echo "Testing"
	@GO15VENDOREXPERIMENT=1 godep go test -race -v $(PGM_PATH)
	@GO15VENDOREXPERIMENT=1 godep go test -race -v $(PGM_PATH)/client...

prof:
	@[ ! -d $(PROFDIR) ] && mkdir $(PROFDIR); GO15VENDOREXPERIMENT=1 godep go test -bench . -benchmem -blockprofile $(PROFDIR)/block.out -cover -coverprofile $(PROFDIR)/cover.out -cpuprofile $(PROFDIR)/cpu.out -memprofile $(PROFDIR)/mem.out $(PROFTARGET)

vet:
	@echo "Vetting"
	@GO15VENDOREXPERIMENT=1 godep go tool vet --all -shadow ./*.go
	@GO15VENDOREXPERIMENT=1 godep go tool vet -all -shadow ./client
	@GO15VENDOREXPERIMENT=1 godep go tool vet -all -shadow ./client/http

dep-save:
	@echo "Run godep save"
	@GO15VENDOREXPERIMENT=1 godep save -v $(SAVE_TARGET)

dep-saved-build: dep-save build

lint:
	@echo "Linting"
	@GO15VENDOREXPERIMENT=1 ${GOBIN}/golint $(PGM_PATH)
	@GO15VENDOREXPERIMENT=1 ${GOBIN}/golint $(PGM_PATH)/client
	@GO15VENDOREXPERIMENT=1 ${GOBIN}/golint $(PGM_PATH)/client/http

release:
	@echo "Releasing"
	@./scripts/release.sh

publish: release
	@echo "Publishing releases to github"
	@./scripts/publish.sh

formula:
	@echo "Generating formula"
	@./scripts/publish.sh formula-only
