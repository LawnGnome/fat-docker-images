GO := go
GOPATH := $(shell pwd)

export GOPATH

all: bin/update

clean:
	rm -rf bin pkg src/github.com src/gopkg.in

bin/update: $(wildcard src/fat/*.go) $(wildcard src/update/*.go)
	$(GO) get fat update
	$(GO) install update
