default: test

help:
	@echo ${MAKE} test
	@echo ${MAKE} xchains_test_bin
	@echo ${MAKE} all

all: xchains_test_bin xchains_racetest_bin

test:
	go test -v -a ./...
race:
	go test -v -a -race ./...

.PHONY += test
.PHONY += race


xchains_test_bin:
	go test -v -a -c -o $@ ./...
xchains_racetest_bin:
	go test -v -a -c -race -o $@ ./...

.PHONY += xchains_test_bin
