tidy:
	go fmt ./...
	go mod tidy

run: build
	./dist/ssm-k8s-injector_darwin_arm64/ssm-k8s-injector

build:
	goreleaser build --snapshot --rm-dist --single-target

clean:
	rm -rf ./dist	
