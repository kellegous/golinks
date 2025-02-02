all: bin/golinks

bin/golinks: main.go $(shell find internal -name '*.go')
	go build -o bin/golinks .

test:
	go test ./...

clean:
	rm -rf bin