GO ?= go
GODEP ?= godep
BINNAME := gat
PGMPKGPATH := .
TESTTARGET := ./...
PROFDIR := ./.profile
PROFTARGET := ./client

all: depbuild

proftest:
	[ ! -d $(PROFDIR) ] && mkdir $(PROFDIR); $(GO) test -bench . -benchmem -blockprofile $(PROFDIR)/block.out -cover -coverprofile $(PROFDIR)/cover.out -cpuprofile $(PROFDIR)/cpu.out -memprofile $(PROFDIR)/mem.out $(PROFTARGET)

depbuild: depsave
	$(GODEP) $(GO) build -o $(GOBIN)/$(BINNAME) $(PGMPKGPATH)

deptest: depsave
	$(GODEP) $(GO) test -v $(TESTTARGET)

depsave:
	$(GODEP) save ./...
