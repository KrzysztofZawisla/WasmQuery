package shortquery

import (
	"syscall/js"
)

var WasmQueryAddClass js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return this.Get("classList")
	}
	var selector string = args[0].String()
	if selector == "undefined" || selector == "" {
		return this.Get("classList")
	}
	this.Get("classList").Call("add", selector)
	return this.Get("classList")
})
