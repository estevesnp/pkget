build:
	@go build -o bin/pkgo .

run: build
	@./bin/pkgo

test:
	@go test -v ./...
