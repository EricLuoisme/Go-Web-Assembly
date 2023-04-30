package assembly

import (
	"github.com/bytecodealliance/wasmtime-go"
)

// createWasmVM is for staring the web assembly engine
func createWasmVM(code []byte) {
	engine := wasmtime.NewEngine()
	module, _ := wasmtime.NewModule(engine, code)
	store := wasmtime.NewStore(engine)
	linker := wasmtime.NewLinker(engine)
	inst, _ := linker.Instantiate(store, module)
	_ = inst
}
