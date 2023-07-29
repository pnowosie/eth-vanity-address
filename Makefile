OUTFILE := vanity.log

GITCOMMIT := $(shell git rev-parse HEAD)
VERSION := 0.0.2

eth-vanity-address: *.go go.mod
	go build -v -o $@ \
	-ldflags="-X main.GitCommit=${GITCOMMIT} -X main.Version=${VERSION}-dev"

run: eth-vanity-address
	@echo "Started with make-added params: [ ${ARGS} ]"
	./eth-vanity-address ${ARGS} 2>> ${OUTFILE}

local-build:
	goreleaser build --snapshot --single-target --rm-dist

release-tag:
	git tag -a "v${VERSION}"
	git push origin v${VERSION}


.PHONY: \
  run \
  local-build \
  release-tag
