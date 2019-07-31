package shortquery

import (
	"syscall/js"
)

var WasmQueryHasClass js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return this.Get("classList")
	}
	var selector string = args[0].String()
	return this.Get("classList").Call("contains", selector)
})
