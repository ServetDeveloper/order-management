build:
	@go build -o bin/order-management cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/order-management

