APP_NAME=crawler
CMD_PATH=./cmd/crawler

.PHONY: build run test fmt tidy clean

build:
	go build -o $(APP_NAME) $(CMD_PATH)

run:
	go run $(CMD_PATH)/main.go

test:
	go test ./...

fmt:
	go fmt ./...

tidy:
	go mod tidy

clean:
	rm -f $(APP_NAME)
