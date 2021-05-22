test:
	go test ./...

lint:
	golangci-lint run

build:
	go build --tags "json1" -o bin/server cmd/server/main.go

run-local-no-build:
	./bin/server --dotenv .env.local

run-local: build run-local-no-build

run-local-ro: build
	./bin/server --dotenv .env.local.ro

gen:
	go generate ./...

clean:
	find . -name "*.gen.*" -type f -delete
	find . -name "*.db" -type f -delete
