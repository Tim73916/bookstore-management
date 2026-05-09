.PHONY: run build clean

run:
	go run ./cmd/main/main.go

build:
	go build -o bookstore ./cmd/main/main.go

clean:
	rm -rf bookstore