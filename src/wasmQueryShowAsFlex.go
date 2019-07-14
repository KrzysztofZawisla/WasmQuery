package main

import "syscall/js"

var wasmQueryShowAsFlex js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	this.Get("style").Set("display", "flex")
	return this.Get("style").Get("display")
})
