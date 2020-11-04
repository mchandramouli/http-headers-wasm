package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"testing"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxytest"
)

func TestHttpHeaders_RequestIdInjectedInResponseHeaders(t *testing.T) {
	opt := proxytest.NewEmulatorOption().
		WithNewHttpContext(newContext)
	host := proxytest.NewHostEmulator(opt)
	defer host.Done()
	id := host.HttpFilterInitContext()

	reqHs := [][2]string{{"x-request-id", "request-1"}, {"User-Agent", "curl/7.64.1"}}
	resHs := [][2]string{{"content-type", "text/plain"}}
	host.HttpFilterPutRequestHeaders(id, reqHs)  // call OnHttpRequestHeaders
	host.HttpFilterPutResponseHeaders(id, resHs) // call OnHttpRequestHeaders

	// x-request-id injected into response header
	modifiedResHs := host.HttpFilterGetResponseHeaders(id)
	assert.Equal(t, 2, len(modifiedResHs))
	assert.Equal(t, "x-request-id", modifiedResHs[1][0])
	assert.Equal(t, "request-1", modifiedResHs[1][1])
}

func TestHttpHeaders_RandomIdInjectedInRequestHeaders(t *testing.T) {
	opt := proxytest.NewEmulatorOption().
		WithNewHttpContext(newContext)
	host := proxytest.NewHostEmulator(opt)
	defer host.Done()
	id := host.HttpFilterInitContext()

	reqHs := [][2]string{{"x-request-id", "request-1"}, {"User-Agent", "curl/7.64.1"}}
	host.HttpFilterPutRequestHeaders(id, reqHs)  // call OnHttpRequestHeaders

	// x-request-id injected into response header
	modifiedReqHs := host.HttpFilterGetRequestHeaders(id)
	assert.Equal(t, 3, len(modifiedReqHs))
	assert.Equal(t, "x-context-key", modifiedReqHs[2][0])
}

func TestHttpHeaders_OnHttpStreamDoneLogsRequestId(t *testing.T) {
	opt := proxytest.NewEmulatorOption().
		WithNewHttpContext(newContext)
	host := proxytest.NewHostEmulator(opt)
	defer host.Done()
	id := host.HttpFilterInitContext()

	reqHs := [][2]string{{"x-request-id", "request-1"}, {"User-Agent", "curl/7.64.1"}}
	host.HttpFilterPutRequestHeaders(id, reqHs)  // call OnHttpRequestHeaders
	host.HttpFilterCompleteHttpStream(id)

	// test logs
	logs := host.GetLogs(types.LogLevelInfo)
	logsLen := len(logs)
	assert.Greater(t, logsLen, 1)
	assert.Equal(t, fmt.Sprintf("request-1 : %d finished", id), logs[len(logs)-1])
}