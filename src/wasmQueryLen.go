package main

import "syscall/js"

var wasmQueryLen js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	return this.Get("length")
})
