package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	proxywasm.SetNewHttpContext(newContext)
}

type wrappedHttpContext struct {
	contextID uint32
	requestId string
	proxywasm.DefaultHttpContext
}

func newContext(rootContextID, contextID uint32) proxywasm.HttpContext {
	httpCtx := &wrappedHttpContext{contextID: contextID}
	proxywasm.LogInfof("newContext --> %v", httpCtx.contextID)
	return httpCtx
}

// override
func (ctx *wrappedHttpContext) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	value, err := proxywasm.GetHttpRequestHeader("x-request-id")
	if err != nil {
		proxywasm.LogCriticalf("failed to get x-request-id: %v", err)
	} else {
		ctx.requestId = value
	}

	//injecting headers
	addHeader("x-context-key", "some-data-goes-here-" +
		strconv.FormatUint(rand.New(rand.NewSource(time.Now().UnixNano())).Uint64(), 16),
		proxywasm.AddHttpRequestHeader)

	printHeaders(proxywasm.GetHttpRequestHeaders, "request")

	return types.ActionContinue
}

// override
func (ctx *wrappedHttpContext) OnHttpResponseHeaders(numHeaders int, endOfStream bool) types.Action {
	if ctx.requestId != "" {
		addHeader("x-request-id", ctx.requestId, proxywasm.AddHttpResponseHeader)
	}

	printHeaders(proxywasm.GetHttpResponseHeaders, "response")

	return types.ActionContinue
}

// override
func (ctx *wrappedHttpContext) OnHttpStreamDone() {
	proxywasm.LogInfof("%v : %v finished", ctx.requestId, ctx.contextID)
}

func printHeaders(getHeaders func() ([][2]string, error), name string) {
	hs, err := getHeaders()
	if err != nil {
		proxywasm.LogCriticalf("failed to %s headers: %v", name, err)
	}

	for _, h := range hs {
		proxywasm.LogInfof("%s [%s = %s]", name, h[0], h[1])
	}
}

func addHeader(name string, value string, setHdr func(string, string) error) {
	proxywasm.LogInfof("add header [%s = %s]", name, value)
	err := setHdr(name, value)
	if err != nil {
		panic("Error setting http header: " + err.Error())
	}
}
