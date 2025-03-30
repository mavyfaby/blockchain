BIN=blockchain

api:
	go mod tidy
	go build -ldflags "-s -w" -o bin/$(BIN)-api cmd/api/main.go
	./bin/$(BIN)-api

blockchain:
	go mod tidy
	go build -ldflags "-s -w" -o bin/$(BIN) cmd/blockchain/main.go
	./bin/$(BIN)

clean:
	go clean
	rm -rf bin
