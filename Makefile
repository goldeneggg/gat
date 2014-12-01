GO ?= go
GODEP ?= godep
GOLINT ?= golint
BINNAME := gat
PGMPKGPATH := .
TESTTARGET := ./...
SAVETARGET := ./...
PROFDIR := ./.profile
PROFTARGET := ./client
LINTTARGET := ./...

all: depbuild

proftest:
	[ ! -d $(PROFDIR) ] && mkdir $(PROFDIR); $(GO) test -bench . -benchmem -blockprofile $(PROFDIR)/block.out -cover -coverprofile $(PROFDIR)/cover.out -cpuprofile $(PROFDIR)/cpu.out -memprofile $(PROFDIR)/mem.out $(PROFTARGET)

depbuild: depsave
	$(GODEP) $(GO) build -o $(GOBIN)/$(BINNAME) $(PGMPKGPATH)

deptest: depvet
	$(GODEP) $(GO) test -race -v $(TESTTARGET)

depvet: depsave
	$(GODEP) $(GO) vet -n $(TESTTARGET)

depsave:
	$(GODEP) save $(SAVETARGET)

lint:
	$(GOLINT) $(LINTTARGET)
