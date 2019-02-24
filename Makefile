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

.PHONY: vendor
vendor:
	@GO111MODULE=on go mod vendor

vendor-build:
	@echo "Building bin/$(NAME) using vendor libraries"
	@go build -mod vendor -o bin/$(NAME) $(PGM_PATH)

lint-travis:
	@travis lint .travis.yml

test-goreleaser:
	@echo "Testing goreleaser"
	@goreleaser release --snapshot --skip-publish --rm-dist

# NOTE: require commands for .travis.yml "before_install" task
# - curl -sL -o goreleaser.tar.gz https://github.com/goreleaser/goreleaser/releases/download/v0.101.0/goreleaser_Linux_x86_64.tar.gz
# - tar -zxf goreleaser.tar.gz
test-goreleaser-on-ci:
	@echo "Testing goreleaser (on CI)"
	@./goreleaser release --snapshot --skip-publish --rm-dist

# .PHONY: docker-test-goreleaser
# docker-test-goreleaser:
# 	@docker run --rm --privileged \
# 		-v $$PWD:/go/src/github.com/goldeneggg/gat \
# 		-v /var/run/docker.sock:/var/run/docker.sock \
# 		-w /go/src/github.com/goldeneggg/gat \
# 		-e GITHUB_TOKEN \
# 		-e GO111MODULE \
# 		goreleaser/goreleaser release --snapshot --skip-publish --rm-dist

.PHONY: ci
ci: test lint

# release process for manual operation
# - (merge pull request)
# - git checkout master && git pull --rebase origin master
# - git tag -a vX.X.X -m 'tag comment'
# - git push --tags
# - make goreleaser
.PHONY: goreleaser
goreleaser:
	@goreleaser release --debug --rm-dist
