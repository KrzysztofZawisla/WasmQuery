package main

import "syscall/js"

var wasmQueryWidth js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return this.Get("style").Get("width")
	}
	value := args[0]
	this.Get("style").Set("width", value)
	return this.Get("style").Get("width")
})
