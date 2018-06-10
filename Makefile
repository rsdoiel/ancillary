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

cmd/ancillary-demo/assets.go: website .FORCE
	pkgassets -o cmd/ancillary-demo/assets.go -p main -ext=".html" Demo "."
	git add cmd/ancillary-demo/assets.go

bin/ancillary-demo$(EXT): ancillary.go cmd/ancillary-demo/ancillary-demo.go cmd/ancillary-demo/assets.go
	go build -o bin/ancillary-demo$(EXT) cmd/ancillary-demo/ancillary-demo.go cmd/ancillary-demo/assets.go

install: 
	env GOBIN=$(GOPATH)/bin go install cmd/ancillary-demo/ancillary-demo.go cmd/ancillary-demo/assets.go

website: page.tmpl README.md LICENSE css/site.css
	bash mk-website.bash

clean: 
	if [ -f index.html ]; then rm *.html; fi
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi

dist/linux-amd64:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/ancillary-demo cmd/ancillary-demo/ancillary-demo.go cmd/ancillary-demo/assets.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-amd64.zip README.md LICENSE bin/*
	rm -fR dist/bin

dist/windows-amd64:
	mkdir -p dist/bin
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/ancillary.exe cmd/ancillary-demo/ancillary-demo.go cmd/ancillary-demo/assets.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-windows-amd64.zip README.md LICENSE bin/*
	rm -fR dist/bin

dist/macosx-amd64:
	mkdir -p dist/bin
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/ancillary cmd/ancillary-demo/ancillary-demo.go cmd/ancillary-demo/assets.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macosx-amd64.zip README.md LICENSE bin/*
	rm -fR dist/bin

dist/raspbian-arm7:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/ancillary cmd/ancillary-demo/ancillary-demo.go cmd/ancillary-demo/assets.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-arm7.zip README.md LICENSE bin/*
	rm -fR dist/bin

distribute_docs:
	if [ -d dist ]; then rm -fR dist; fi
	mkdir -p dist
	cp -v README.md dist/
	cp -v LICENSE dist/
	bash package-versions.bash > dist/package-versions.txt

update_version:
	./update_version.py --yes

release: clean ancillary.go distribute_docs dist/linux-amd64 dist/windows-amd64 dist/macosx-amd64 dist/raspbian-arm7

status:
	git status

save:
	if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

publish:
	bash mk-website.bash
	bash publish.bash

.FORCE:
