default: test

help:
	@echo ${MAKE} test
	@echo ${MAKE} xchains_test_bin
	@echo ${MAKE} all

all: xchains_test_bin xchains_racetest_bin

test:
	go test -v -count=1 ./...
race:
	go test -v -count=1 -race ./...

.PHONY += test
.PHONY += race


xchains_test_bin:
	go test -v -count=1 -c -o $@ ./...
xchains_racetest_bin:
	go test -v -count=1 -c -race -o $@ ./...

.PHONY += xchains_test_bin
.PHONY += xchains_racetest_bin

clean:
	${RM} xchains_test_bin xchains_racetest_bin
