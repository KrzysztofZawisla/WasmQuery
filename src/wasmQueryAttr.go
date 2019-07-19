package main

import (
	"syscall/js"
)

var wasmQueryAttr js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return nil
	}
	var selector string = args[0].String()
	if len(args) >= 1 {
		return this.Call("getAttribute", selector)
	}
	return nil
})
