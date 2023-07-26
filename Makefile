OUTFILE := vanity.log

.PHONY: run
run: eth-vanity-address
	@echo "Started with make-added params: [ ${ARGS} ]"
	./eth-vanity-address ${ARGS} 2>> ${OUTFILE}

eth-vanity-address: *.go go.mod
	go build -o $@ .
