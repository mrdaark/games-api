.PHONY: build
build:
	go build -o ./dist/ -v ./cmd/main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL != build
