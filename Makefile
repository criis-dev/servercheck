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

install:
	sudo mkdir -p /etc/servercheck
	sudo cp dist/servercheck-linux-amd64 /usr/local/bin/servercheck
	sudo cp configs/servers.yaml /etc/servercheck/servers.yaml
	sudo chmod +x /usr/local/bin/servercheck