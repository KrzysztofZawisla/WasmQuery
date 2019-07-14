package main

import "syscall/js"

var wasmQueryVal js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return this.Get("value")
	}
	value := args[0]
	this.Set("value", value)
	return this.Get("value")
})
