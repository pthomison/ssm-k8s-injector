tidy:
	go fmt ./...
	go mod tidy

run: build
	./dist/ssm-k8s-injector_darwin_arm64/ssm-k8s-injector \
		--ssm-parameter "/placeholder"\
		--k8s-secret "test-secret"
	./dist/ssm-k8s-injector_darwin_arm64/ssm-k8s-injector \
		--ssm-parameter "/placeholder1"\
		--k8s-secret "test-secret"
	./dist/ssm-k8s-injector_darwin_arm64/ssm-k8s-injector \
		--ssm-parameter "/placeholder2"\
		--k8s-secret "test-secret"
	kubectl delete secret test-secret -n default

build:
	goreleaser build --snapshot --rm-dist --single-target

clean:
	rm -rf ./dist

release:
	goreleaser release --rm-dist
