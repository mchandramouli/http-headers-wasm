.PHONY: clean test build build-docker

clean:
	rm -rf *.wasm

build: clean
	tinygo build -o ./http-headers.wasm -scheduler=none -target=wasi -wasm-abi=generic ./main.go

build-docker:
	docker run -it -w /tmp/http-headers-wasm -v $(shell pwd):/tmp/http-headers-wasm tinygo/tinygo-dev:latest make build
