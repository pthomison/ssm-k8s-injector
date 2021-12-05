tidy:
	go fmt ./...
	go mod tidy

run: build
	./dist/ssm-k8s-injector_darwin_arm64/ssm-k8s-injector \
		--ssm-parameter "/placeholder"\
		--secret-name "test-secret" \
		--secret-key "placeholderKeyTest"
	./dist/ssm-k8s-injector_darwin_arm64/ssm-k8s-injector \
		--ssm-parameter "/placeholder1"\
		--secret-name "test-secret"
	./dist/ssm-k8s-injector_darwin_arm64/ssm-k8s-injector \
		--ssm-parameter "/placeholder2"\
		--secret-name "test-secret" \
		--secret-key "placeholderKeyTwoTest"
# 	kubectl delete secret test-secret -n default

build:
	goreleaser build --snapshot --rm-dist --single-target

clean:
	rm -rf ./dist

release:
	goreleaser release --rm-dist
