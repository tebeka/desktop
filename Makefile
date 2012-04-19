export GOPATH := $(shell dirname $(shell dirname $(PWD)))
PACKAGE := desktop

all:
	go build $(PACKAGE)

test:
	go test -v $(PACKAGE)

fix:
	go fix $(PACKAGE)

doc:
	go doc $(PACKAGE)

install:
	go install $(PACKAGE)

push:
	hg push
	hg push git+ssh://git@github.com/tebeka/desktop.git

README.html: README.rst
	rst2html $< > $@

.PHONY: all test install fix doc
