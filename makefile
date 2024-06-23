run:
	go run cmd/main.go

lint:
	golangci-lint run

test:
	go test -v cover ./...