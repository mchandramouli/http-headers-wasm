.PHONY: clean test build build.docker

build:
	tinygo build -o ./hello_world.wasm -scheduler=none -target=wasi -wasm-abi=generic ./hello_world.go

build.docker:
	docker run -it -w /tmp/hello_world_wasm -v $(shell pwd):/tmp/hello_world_wasm tinygo/tinygo-dev:latest \
		tinygo build -o /tmp/hello_world_wasm/hello_world.wasm -scheduler=none -target=wasi \
		-wasm-abi=generic /tmp/hello_world_wasm/hello_world.go
