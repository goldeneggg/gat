GO ?= go
GODEP ?= godep
BINNAME := gat
PGMPKGPATH := .
TESTTARGET := ./...
PROFDIR := ./.profile
PROFTARGET := ./client

all: depbuild

build: getdeps
	$(GO) build -o $(GOBIN)/$(BINNAME) $(PGMPKGPATH)

test:
	$(GO) test -v $(TESTTARGET)

proftest:
	[ ! -d $(PROFDIR) ] && mkdir $(PROFDIR); $(GO) test -bench . -benchmem -blockprofile $(PROFDIR)/block.out -cover -coverprofile $(PROFDIR)/cover.out -cpuprofile $(PROFDIR)/cpu.out -memprofile $(PROFDIR)/mem.out $(PROFTARGET)


# Following targets using "godep"
depbuild: depsave
	$(GODEP) $(GO) build -o $(GOBIN)/$(BINNAME) $(PGMPKGPATH)

deptest: depsave
	$(GODEP) $(GO) test -v $(TESTTARGET)

depsave:
	$(GODEP) save ./...
