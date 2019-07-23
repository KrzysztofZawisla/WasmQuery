package core

import (
	"syscall/js"

	"github.com/KrzysztofZawisla/WasmQuery/shortquery"
	timequery "github.com/KrzysztofZawisla/WasmQuery/timeQuery/wasmQueryTime"
)

/* WasmQueryRegisterLibrary is a function to register library. */
var WasmQueryRegisterLibrary js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	var symbol string = "$"
	if len(args) >= 1 {
		if args[0].Type().String() == "object" {
			var outputArray []interface{}
			for i := 0; i < args[0].Length(); i++ {
				js.Global().Set(args[0].Index(i).String(), shortquery.WasmQuery)
				js.Global().Get(args[0].Index(i).String()).Set("disableLibrary", WasmQueryDisableLibrary)
				js.Global().Get(args[0].Index(i).String()).Set("timeNow", timequery.WasmQueryTimeNow)
				outputArray = append(outputArray, args[0].Index(i))
			}
			return outputArray
		}
		symbol = args[0].String()
	}
	js.Global().Set(symbol, shortquery.WasmQuery)
	js.Global().Get(symbol).Set("disableLibrary", WasmQueryDisableLibrary)
	js.Global().Get(symbol).Set("timeNow", timequery.WasmQueryTimeNow)
	return symbol
})
