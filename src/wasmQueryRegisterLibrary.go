package main

import (
	"syscall/js"
)

var wasmQueryRegisterLibrary js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	var symbol string = "$"
	if len(args) >= 1 {
		symbol = args[0].String()
	} 
	js.Global().Set(symbol, wasmQuery)
	js.Global().Get(symbol).Set("disableLibrary", wasmQueryDisableLibrary)
}
