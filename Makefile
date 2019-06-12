GO = GO111MODULE=on GOFLAGS=-mod=vendor go
VERSION = 0.1.0
BINARY = super-hacker
BINDIR = bin
LDFLAGS = -ldflags "-X main.gitSHA=$(shell git rev-parse HEAD) -X main.version=$(VERSION) -X main.name=$(BINARY)"

$(BINDIR)/$(BINARY): $(BINDIR)
	$(GO) build -v -o $(BINDIR)/$(BINARY) $(LDFLAGS)

$(BINDIR):
	mkdir $(BINDIR) 

.PHONY: deps
deps:
	$(GO) mod download
	$(GO) mod vendor

.PHONY: test
test:
	$(GO) test -v -cover ./...

.PHONY: clean
clean:
	$(GO) clean
	rm -f $(BINDIR)/$(BINARY)
