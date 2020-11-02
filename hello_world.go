package hello

import (
	"math/rand"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
)

const tickMilliseconds uint32 = 1000

func main() {
	proxywasm.SetNewRootContext(newHelloWorld)
}

type helloWorld struct {
	proxywasm.DefaultRootContext
	contextID uint32
}

func newHelloWorld(contextID uint32) proxywasm.RootContext {
	return &helloWorld{contextID: contextID}
}

// override
func (ctx *helloWorld) OnVMStart(vmConfigurationSize int) bool {
	rand.Seed(proxywasm.GetCurrentTime())

	proxywasm.LogInfo("proxy_on_vm_start from Go!")
	if err := proxywasm.SetTickPeriodMilliSeconds(tickMilliseconds); err != nil {
		proxywasm.LogCriticalf("failed to set tick period: %v", err)
	}

	return true
}

// override
func (ctx *helloWorld) OnTick() {
	t := proxywasm.GetCurrentTime()
	proxywasm.LogInfof("It's %d: random value: %d", t, rand.Uint64())
}
