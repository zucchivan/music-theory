all: test
TESTS := expr unrecognised

.PHONY: test fmt

fmt:
	go fmt ./...

test:
	go get -v ./... && go test ./...
