run:
	go run ./cmd/todo

build:
	go build -o todo ./cmd/todo

clean:
	rm -f todo

test:
	go test ./...
