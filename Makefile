BINARY=servercheck
DIST=dist

build:
	go build -o $(BINARY) ./cmd/servercheck

clean:
	rm -f $(BINARY)

run:
	go run ./cmd/servercheck

test:
	go test ./...

fmt:
	gofmt -w .

release:
	mkdir -p $(DIST)
	GOOS=linux GOARCH=amd64 go build -o $(DIST)/servercheck-linux-amd64 ./cmd/servercheck