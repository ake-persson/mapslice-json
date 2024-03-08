all:	test

deps:
	go install golang.org/x/lint/golint@latest
	go install github.com/kisielk/errcheck@latest
	go install github.com/client9/misspell/cmd/misspell@latest
	go install github.com/gordonklaus/ineffassign@latest
	go get -u github.com/fzipp/gocyclo@latest
	go install github.com/fzipp/gocyclo/cmd/gocyclo@latest


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
	go test -race -v

coverage:
	go test -v -covermode=count -coverprofile=coverage.out
	go tool cover -html=coverage.out

.PHONY: deps clean fmt test coverage
