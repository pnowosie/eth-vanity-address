OUTFILE := vanity.log

GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell TZ=UTC0 git show -s --date='format-local:%Y-%m-%dT%H:%M:%S+00' --format="%cd")
VERSION := 0.0.1

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.Version=$(VERSION)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

eth-vanity-address: *.go go.mod
	go build -v $(LDFLAGS) -o $@ .

run: eth-vanity-address
	@echo "Started with make-added params: [ ${ARGS} ]"
	./eth-vanity-address ${ARGS} 2>> ${OUTFILE}

local-build:
	goreleaser build --snapshot --single-target --rm-dist

release-tag:
	git tag -a "v${VERSION}"
	git push origin v${VERSION}

release-publish: release-tag
	goreleaser release --clean

.PHONY: \
  run \
  local-build \
  release-tag \
  release-publish
