-include .env

export PORT
export APP_ENV
export MONGODB_URI
export MONGODB_DATABASE
export MONGODB_COLLECTION

.PHONY: run build tidy clean

run:
	go run ./cmd/server

build:
	mkdir -p bin
	go build -o bin/server ./cmd/server

tidy:
	go mod tidy

clean:
	rm -rf bin