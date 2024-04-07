all: build test
build:
	go build
lint:
	go vet ./...
format:
	gofmt -w .
	golangci-lint run
test:
	go test