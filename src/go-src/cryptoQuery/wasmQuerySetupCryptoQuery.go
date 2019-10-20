package cryptoquery

import (
	"syscall/js"
)

// WasmQueryFunctionsStorage stores all functions
var WasmQueryFunctionsStorage = make(map[string]js.Func)

// SetUpFunctionsToVariables register function to case in map
func SetUpFunctionsToVariables() {
	WasmQueryFunctionsStorage["WasmQueryKeccak256"] = js.FuncOf(WasmQueryKeccak256)
	WasmQueryFunctionsStorage["WasmQueryKeccak512"] = js.FuncOf(WasmQueryKeccak512)
	WasmQueryFunctionsStorage["WasmQueryMD4"] = js.FuncOf(WasmQueryMD4)
	WasmQueryFunctionsStorage["WasmQueryMD5"] = js.FuncOf(WasmQueryMD5)
	WasmQueryFunctionsStorage["WasmQuerySHA1"] = js.FuncOf(WasmQuerySHA1)
	WasmQueryFunctionsStorage["WasmQuerySHA3_224"] = js.FuncOf(WasmQuerySHA3_224)
	WasmQueryFunctionsStorage["WasmQuerySHA3_256"] = js.FuncOf(WasmQuerySHA3_256)
	WasmQueryFunctionsStorage["WasmQuerySHA3_384"] = js.FuncOf(WasmQuerySHA3_384)
	WasmQueryFunctionsStorage["WasmQuerySHA3_512"] = js.FuncOf(WasmQuerySHA3_512)
	WasmQueryFunctionsStorage["WasmQuerySHA224"] = js.FuncOf(WasmQuerySHA224)
	WasmQueryFunctionsStorage["WasmQuerySHA256"] = js.FuncOf(WasmQuerySHA256)
	WasmQueryFunctionsStorage["WasmQuerySHA384"] = js.FuncOf(WasmQuerySHA384)
	WasmQueryFunctionsStorage["WasmQuerySHA512"] = js.FuncOf(WasmQuerySHA512)
	WasmQueryFunctionsStorage["WasmQuerySHA512_224"] = js.FuncOf(WasmQuerySHA512_224)
	WasmQueryFunctionsStorage["WasmQuerySHA512_256"] = js.FuncOf(WasmQuerySHA512_256)
}
