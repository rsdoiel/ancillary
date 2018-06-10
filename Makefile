#
# Simple Makefile for conviently testing, building and deploying 
# experimental projects.
#
PROJECT = ancillary

VERSION = $(shell grep -m 1 'Version =' $(PROJECT).go | cut -d\`  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

PKGASSETS = $(shell which pkgassets)

PROJECT_LIST = ancillary

OS = $(shell uname)

EXT = 
ifeq ($(OS), Windows)
	EXT = .exe
endif

ancillary-demo$(EXT): bin/ancillary-demo$(EXT)

cmd/ancillary-demo/css.go: .FORCE
	pkgassets -o cmd/ancillary-demo/css.go -p main -ext=".css" CSS "."
	git add cmd/ancillary-demo/css.go

cmd/ancillary-demo/html.go: website .FORCE
	pkgassets -o cmd/ancillary-demo/html.go -p main -ext=".html" HTML "."
	git add cmd/ancillary-demo/html.go

bin/ancillary-demo$(EXT): ancillary.go cmd/ancillary-demo/ancillary-demo.go cmd/ancillary-demo/html.go cmd/ancillary-demo/css.go
	go build -o bin/ancillary-demo$(EXT) cmd/ancillary-demo/ancillary-demo.go cmd/ancillary-demo/html.go cmd/ancillary-demo/css.go

install: .FORCE
	env GOBIN=$(GOPATH)/bin go install cmd/ancillary-demo/ancillary-demo.go cmd/ancillary-demo/html.go cmd/ancillary-demo/css.go

website: page.tmpl README.md LICENSE css/site.css
	bash mk-website.bash

clean: 
	if [ -f index.html ]; then rm *.html; fi
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi

update_version:
	./update_version.py --yes

status:
	git status

save: website
	if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

publish:
	bash mk-website.bash
	bash publish.bash

.FORCE:
