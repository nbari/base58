.PHONY: test bench cover clean

GO ?= go

test:
	${GO} test -v

bench:
	${GO} test -bench=.

cover:
	${GO} test -cover && \
	${GO} test -coverprofile=coverage.out  && \
	${GO} tool cover -html=coverage.out

clean:
	rm *.out
