package main

import "syscall/js"

var wasmQueryShowAsInlineBlock js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	this.Get("style").Set("display", "inline-block")
	return this.Get("style").Get("display")
})
