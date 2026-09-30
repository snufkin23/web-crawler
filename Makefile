.PHONY: run build fmt test check clean

APP_NAME=build
MAIN_FILE=cmd/main.go
BIN=bin/$(APP_NAME)

run:
	go run $(MAIN_FILE)

build:
	@mkdir -p bin
	go build -o $(BIN) $(MAIN_FILE)

fmt:
	go fmt ./...

test:
	go test ./...

check:
	go vet ./...

clean:
	rm -rf bin
	go clean -testcache