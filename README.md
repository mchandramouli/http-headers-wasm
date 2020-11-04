## http-headers-wasm

Sample [envoy](https://www.envoyproxy.io/) http filter using a [wasm](https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/http/wasm/v3/wasm.proto#extensions-filters-http-wasm-v3-wasm) module

This code uses [proxy-wasm-go-sdk](https://github.com/tetratelabs/proxy-wasm-go-sdk) and [tinygo](https://tinygo.org/) to build the wasm module 

### Building the module

Make will launch a docker container with `tinygo` to build and run unit tests

```bash
$ make 
```    

### Testing locally

One can test this locally using one of the two methods

#### Using docker

```bash        
$ make run-with-docker
```   

#### Using docker-compose

```bash
$ docker-compose -f test/docker-compose.yaml --project-directory $(pwd) up
```      

Both of these expose the proxy in port `8000`. Try `curl http://localhost:8000 -v` and watch for response headers and proxy logs.

