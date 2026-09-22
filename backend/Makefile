.PHONY: all build run run-worker seed test tidy clean

all: build

build:
	@mkdir -p bin
	go build -o bin/api ./cmd/api
	go build -o bin/worker ./cmd/worker

run:
	go run ./cmd/api

run-worker:
	go run ./cmd/worker

seed:
	go run ./cmd/seed

test:
	go test -v ./...

tidy:
	go mod tidy

clean:
	rm -rf bin/
