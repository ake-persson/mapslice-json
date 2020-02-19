all:	test

clean:
	rm -f _example/_example coverage.out

fmt:
	gofmt -w .

test:
	golint -set_exit_status .
	go vet .
	errcheck .
	misspell .
	ineffassign .
	gocyclo -over 15 .
	go test . -v -covermode=atomic

coverage:
	go test -covermode=count -coverprofile=coverage.out
	go tool cover -html=coverage.out

.PHONY: clean fmt test coverage
