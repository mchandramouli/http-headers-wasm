package main

import (
	"math/rand"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"strconv"
)

func main() {
	proxywasm.SetNewHttpContext(newContext)
}

type httpHeaders struct {
	proxywasm.DefaultHttpContext
	contextID uint32
}

func newContext(rootContextID, contextID uint32) proxywasm.HttpContext {
	proxywasm.LogInfof("newContext --> %v", contextID)
	return &httpHeaders{contextID: contextID}
}

func addHeader(name string, value string) {
	err := proxywasm.AddHttpRequestHeader(name, value)
	if err != nil {
		panic("Error setting http header: " + err.Error())
	}
}

// override
func (ctx *httpHeaders) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	hs, err := proxywasm.GetHttpRequestHeaders()
	if err != nil {
		proxywasm.LogCriticalf("failed to get request headers: %v", err)
	}

	for _, h := range hs {
		proxywasm.LogInfof("request header --> %s: %s", h[0], h[1])
	}

	//injecting headers
	addHeader("x-request-id", strconv.FormatUint(rand.Uint64(), 10))

	return types.ActionContinue
}
