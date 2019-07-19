package main

import (
	"syscall/js"
	"time"
)

var wasmQueryTimeNow js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	return time.Now().String()
})
