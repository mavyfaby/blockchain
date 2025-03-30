BIN=blockchain

run:
	go mod tidy
	go build -ldflags "-s -w" -o bin/$(BIN) cmd/blockchain/main.go
	./bin/$(BIN)

clean:
	go clean
	rm -rf bin
