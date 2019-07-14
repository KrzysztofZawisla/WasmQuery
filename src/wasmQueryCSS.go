package main

import "syscall/js"

var wasmQueryCSS js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return nil
	}
	var selector string = args[0].String()
	if len(args) == 1 {
		return this.Get("style").Get(selector)
	} else if len(args) >= 2 {
		var value string = args[1].String()
		this.Get("style").Set(selector, value)
		return this.Get("style").Get(selector)
	}
	return nil
})
