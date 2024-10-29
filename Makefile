.PHONY: build run test clean

BINARY_NAME=auction-data-collector

build:
	go build -o $(BINARY_NAME) ./cmd/main.go

run:
	go run ./cmd/main.go

test:
	go test ./...

clean:
	rm -f $(BINARY_NAME)