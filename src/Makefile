.PHONY: test coverage report build clean

setup:
	go install gotest.tools/gotestsum@latest

test:
	gotestsum --format testname -- -cover -v ./...

coverage:
	go test -coverprofile=coverage.out ./...

report: coverage
	go tool cover -html=coverage.out -o coverage.html

build:
	go build ./...

clean:
	rm -f coverage.out coverage.html