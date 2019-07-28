package shortquery

import (
	"syscall/js"
)

var WasmQueryRemoveClass js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return this.Get("classList")
	}
	var selector string = args[0].String()
	this.Get("classList").Call("remove", selector)
	return this.Get("classList")
})
