build:
	go build -o ../../bin/urlshort ./cmd/main.go
test:
	go test ./...