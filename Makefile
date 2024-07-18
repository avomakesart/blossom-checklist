build:
	go build -o blossom-checklist

test: build
	go test -v
