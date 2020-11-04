.PHONY: clean test build build-with-docker run-with-docker

build-with-docker:
	docker run -it -w /tmp/http-headers-wasm -v $(shell pwd):/tmp/http-headers-wasm tinygo/tinygo-dev:latest make build

run-with-docker: build-with-docker
	docker run --entrypoint='/usr/local/bin/envoy' \
		-p 8000:18000 -p 8099:8099 \
		-w /tmp/http-headers-wasm -v $(shell pwd):/tmp/http-headers-wasm envoyproxy/envoy-dev:latest \
		-c /tmp/http-headers-wasm/test/envoy.yaml --concurrency 2

clean:
	rm -rf *.wasm

build: clean test
	tinygo build -o ./http-headers.wasm -scheduler=none -target=wasi -wasm-abi=generic ./main.go

test:
	go test -tags=proxytest
