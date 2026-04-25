run:
	go run cmd/api.go

build:
	mkdir -p bin
	go build -o bin/api cmd/api.go

clean:
	rm -f bin