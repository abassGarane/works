build:
	@go build -o bin/works

run: build
	@./bin/works
test:
	@go test -v ./...
