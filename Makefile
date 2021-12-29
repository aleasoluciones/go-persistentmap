all: clean test build

test:
	go vet ./...
	go clean -testcache
	go test -v -tags integration ./...

build:
	go build -o go-persistentmap sample/sample.go

clean:
	rm -f go-persistentmap


.PHONY: all test build clean
