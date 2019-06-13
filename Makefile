GO      = GO111MODULE=on GOFLAGS=-mod=vendor go
VERSION = 0.1.0
BINARY  = super-hacker
BINDIR  = bin
LDFLAGS = -ldflags "-X main.gitSHA=$(shell git rev-parse HEAD) -X main.version=$(VERSION) -X main.name=$(BINARY)"
OS      = $(shell uname)

$(BINDIR)/$(BINARY): $(BINDIR)
	$(GO) build -v -o $(BINDIR)/$(BINARY) $(LDFLAGS)

$(BINDIR):
	mkdir $(BINDIR) 

install: clean
ifeq ($(OS),Darwin)
	./build.sh darwin $(VERSION)
	cp -f $(BINDIR)/$(BINARY)-darwin /usr/local/$(BINDIR)/$(BINARY)
endif 
ifeq ($(OS),Linux)
	./build.sh linux $(VERSION)
	cp -f $(BINDIR)/$(BINARY)-linux /usr/local/$(BINDIR)/$(BINARY)
endif
ifeq ($(OS),FreeBSD)
	./build.sh freebsd $(VERSION)
	cp -f $(BINDIR)/$(BINARY)-freebsd /usr/local/$(BINDIR)/$(BINARY)
endif
uninstall: 
	rm -f /usr/local/$(BINDIR)/$(BINARY)

.PHONY: deps
deps:
	$(GO) mod download
	$(GO) mod vendor

.PHONY: test
test:
	$(GO) test -v -cover ./...

.PHONY: release
release:
	./build.sh release $(VERSION)

.PHONY: clean
clean:
	$(GO) clean
	rm -f $(BINDIR)/*
