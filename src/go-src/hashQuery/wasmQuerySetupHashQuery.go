package hashquery

import (
	"syscall/js"
)

// WasmQueryFunctionsStorage stores all functions
var WasmQueryFunctionsStorage = make(map[string]js.Func)

// SetUpFunctionsToVariables register function to case in map
func SetUpFunctionsToVariables() {
	WasmQueryFunctionsStorage["WasmQueryAdler32"] = js.FuncOf(WasmQueryAdler32)
	WasmQueryFunctionsStorage["WasmQueryFNV1_32"] = js.FuncOf(WasmQueryFNV1_32)
	WasmQueryFunctionsStorage["WasmQueryFNV1_64"] = js.FuncOf(WasmQueryFNV1_64)
	WasmQueryFunctionsStorage["WasmQueryFNV1_128"] = js.FuncOf(WasmQueryFNV1_128)
	WasmQueryFunctionsStorage["WasmQueryFNV1a32"] = js.FuncOf(WasmQueryFNV1a32)
	WasmQueryFunctionsStorage["WasmQueryFNV1a64"] = js.FuncOf(WasmQueryFNV1a64)
	WasmQueryFunctionsStorage["WasmQueryFNV1a128"] = js.FuncOf(WasmQueryFNV1a128)
}
