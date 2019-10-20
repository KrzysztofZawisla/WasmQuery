package timequery

import (
	"syscall/js"
)

// WasmQueryFunctionsStorage stores all functions
var WasmQueryFunctionsStorage = make(map[string]js.Func)

// SetUpFunctionsToVariables register function to case in map
func SetUpFunctionsToVariables() {
	WasmQueryFunctionsStorage["WasmQueryNow"] = js.FuncOf(WasmQueryNow)
	WasmQueryFunctionsStorage["WasmQueryDay"] = js.FuncOf(WasmQueryDay)
	WasmQueryFunctionsStorage["WasmQueryHour"] = js.FuncOf(WasmQueryHour)
	WasmQueryFunctionsStorage["WasmQueryMinute"] = js.FuncOf(WasmQueryMinute)
	WasmQueryFunctionsStorage["WasmQueryMonth"] = js.FuncOf(WasmQueryMonth)
	WasmQueryFunctionsStorage["WasmQueryYear"] = js.FuncOf(WasmQueryYear)
	WasmQueryFunctionsStorage["WasmQuerySecond"] = js.FuncOf(WasmQuerySecond)
	WasmQueryFunctionsStorage["WasmQueryNanosecond"] = js.FuncOf(WasmQueryNanosecond)
}
