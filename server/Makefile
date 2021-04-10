test:
	go test ./...

build:
	go build --tags "json1" -o bin/server cmd/server/main.go

lint:
	golangci-lint run

run: build
	./bin/server --dotenv .env.local

run-ro: build
	./bin/server --dotenv .env.local.read

run-get-hn-now:
	go run scripts/get-hn-now/main.go
