tidy:
	go fmt ./...
	go mod tidy

build:
	goreleaser build --snapshot

clean:
	rm -rf ./dist	
