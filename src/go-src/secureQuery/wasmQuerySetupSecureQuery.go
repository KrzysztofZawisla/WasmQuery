package securequery

import (
	"syscall/js"
)

// WasmQueryFunctionsStorage stores all functions
var WasmQueryFunctionsStorage = make(map[string]js.Func)

// SetUpFunctionsToVariables register function to case in map
func SetUpFunctionsToVariables() {
	WasmQueryFunctionsStorage["WasmQuerySecureLinks"] = js.FuncOf(WasmQuerySecureLinks)
}
