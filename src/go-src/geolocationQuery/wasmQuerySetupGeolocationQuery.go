package geolocationquery

import (
	"syscall/js"
)

// WasmQueryFunctionsStorage stores all functions
var WasmQueryFunctionsStorage = make(map[string]js.Func)

// SetUpFunctionsToVariables register function to case in map
func SetUpFunctionsToVariables() {
	WasmQueryFunctionsStorage["WasmQueryLatitude"] = js.FuncOf(WasmQueryLatitude)
	WasmQueryFunctionsStorage["WasmQueryLongitude"] = js.FuncOf(WasmQueryLongitude)
}
