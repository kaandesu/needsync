APP_NAME=needsync

build:
	go build -o bin/cli ./cmd/cli
	go build -o bin/api ./cmd/api

run-cli:
	go run ./cmd/cli/main.go sync

run-api:
	go run ./cmd/api/main.go

tidy:
	go mod tidy

clean:
	rm -rf bin/

.PHONY: build run-cli run-api tidy clean
