test:
	go test -v

README.html: README.rst
	rst2html $< > $@

.PHONY: all test install fix doc
