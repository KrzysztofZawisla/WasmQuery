package main

import (
	"syscall/js"
)

var wasmQueryRegisterLibrary js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	var symbol string = "$"
	if len(args) >= 1 {
		if args[0].Type().String() == "object" {
			var outputArray []interface{}
			for i := 0; i < args[0].Length(); i++ {
				js.Global().Set(args[0].Index(i).String(), wasmQuery)
				js.Global().Get(args[0].Index(i).String()).Set("disableLibrary", wasmQueryDisableLibrary)
				js.Global().Get(args[0].Index(i).String()).Set("timeNow", wasmQueryTimeNow)
				outputArray = append(outputArray, args[0].Index(i))
			}
			return outputArray
		}
		symbol = args[0].String()
	}
	js.Global().Set(symbol, wasmQuery)
	js.Global().Get(symbol).Set("disableLibrary", wasmQueryDisableLibrary)
	js.Global().Get(symbol).Set("timeNow", wasmQueryTimeNow)
	return symbol
})
