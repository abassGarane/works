build:
	@go build  -o bin/works cmd/web/main.go
run: build
	@./bin/works
test:
	@go test -v ./...
