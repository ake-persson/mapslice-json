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
	go test . -v -covermode=atomic

cover:
	go test -covermode=count -coverprofile=cover.out fmt 
	go tool cover -html=cover.out

.PHONY: clean fmt test
