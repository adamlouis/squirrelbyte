
COMMIT_SHORT_HASH:=$(shell git log -n 1 --pretty=format:%h)

run:
	go run *.go

use-local-dep:
	go get github.com/adamlouis/squirrelbyte/server@$(COMMIT_SHORT_HASH)
