package main

import "syscall/js"

var wasmQueryShowAsBlock js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	this.Get("style").Set("display", "block")
	return this.Get("style").Get("display")
})
