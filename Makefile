OUTFILE := vanity.log

GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell TZ=UTC0 git show -s --date='format-local:%Y-%m-%dT%H:%M:%S+00' --format="%cd")
VERSION := v0.0.0

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.Version=$(VERSION)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

eth-vanity-address: *.go go.mod
	go build -v $(LDFLAGS) -o $@ .

.PHONY: run
run: eth-vanity-address
	@echo "Started with make-added params: [ ${ARGS} ]"
	./eth-vanity-address ${ARGS} 2>> ${OUTFILE}
