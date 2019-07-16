package main

import "syscall/js"

var wasmQueryHeight js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return this.Get("style").Get("height")
	}
	value := args[0]
	this.Get("style").Set("height", value)
	return this.Get("style").Get("height")
})
