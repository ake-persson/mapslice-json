all:	test

clean:
	rm -f _example/_example

fmt:
	gofmt -w .

test:
	golint -set_exit_status .
	go vet .
	errcheck .
	misspell .
	ineffassign .
	gocyclo -over 15 .

.PHONY: clean fmt test
