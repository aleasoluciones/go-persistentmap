all: clean test build

update_dep:
	go get $(DEP)
	go mod tidy

update_all_deps:
	go get -u all
	go mod tidy

test:
	go vet ./...
	go clean -testcache
	go test -v -tags integration ./...

build:
	go build -o go-persistentmap sample/sample.go

clean:
	rm -f go-persistentmap


.PHONY: all update_dep update_all_deps test build clean
