#!/bin/bash

if [[ ! -f bin/ancillary ]]; then
	echo "Running make before generating updated docs."
	make
fi
bin/ancillary -generate-markdown-docs >"docs/ancillary.md"

