NAME := gat
PGM_PATH := 'github.com/goldeneggg/gat'
SAVE_TARGET := ./...

SRCS = $(shell find . -type f -name '*.go' | \grep -v 'vendor')

.DEFAULT_GOAL := bin/$(NAME)

mod-dl:
	@GO111MODULE=on go mod download

mod-tidy:
	@GO111MODULE=on go mod tidy

mod-golint-install: mod-tidy
	@GO111MODULE=on go install golang.org/x/lint/golint

bin/$(NAME): $(SRCS)
	@echo "Building bin/$(NAME)"
	@go build -o bin/$(NAME) $(PGM_PATH)

.PHONY: test
test:
	@echo "Testing"
	@go test -race -cover -v $$(go list ./... | \grep -v 'vendor')

.PHONY: vet
vet:
	@echo "Vetting"
	@go vet -n -x $$(go list ./... | \grep -v 'vendor')

.PHONY: lint
lint:
	@echo "Linting"
	@golint -set_exit_status $$(go list ./... | \grep -v 'vendor')

.PHONY: ci
ci: bin/$(NAME) test vet lint

.PHONY: vendor
vendor:
	@GO111MODULE=on go mod vendor

vendor-build:
	@echo "Building bin/$(NAME) using vendor libraries"
	@go build -mod vendor -o bin/$(NAME) $(PGM_PATH)

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

# release process for manual operation
# - (merge pull request)
# - git checkout master && git pull --rebase origin master
# - git tag -a vX.X.X -m 'tag comment'
# - git push --tags
# - make goreleaser
.PHONY: goreleaser
goreleaser:
	@goreleaser release --debug --rm-dist
