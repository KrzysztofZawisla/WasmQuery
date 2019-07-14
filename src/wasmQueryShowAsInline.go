package main

import "syscall/js"

var wasmQueryShowAsInline js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	this.Get("style").Set("display", "inline")
	return this.Get("style").Get("display")
})
