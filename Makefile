all:	test

deps:
	go get -u golang.org/x/lint/golint
	go get -u github.com/kisielk/errcheck
	go get -u github.com/client9/misspell/cmd/misspell
	go get -u github.com/gordonklaus/ineffassign
	go get -u github.com/fzipp/gocyclo
	go get -u github.com/fzipp/gocyclo/cmd/gocyclo@latest


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
	go test -v -covermode=count

coverage:
	go test -v -covermode=count -coverprofile=coverage.out
	go tool cover -html=coverage.out

.PHONY: deps clean fmt test coverage
