package main

import "syscall/js"

var wasmQueryHide js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	this.Get("style").Set("display", "none")
	return this.Get("style").Get("display")
})
