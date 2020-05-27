.PHONY: build test

build:
	go build -o steam-gameserver
	
test:
	go test -v ./pkg/...
