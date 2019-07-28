package shortquery

import (
	"syscall/js"
)

/* WasmQueryAttr is a function from shortQuery module. This function set/return attribute for html element. This function can only work on single element. */
var WasmQueryAttr js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return nil
	}
	var selector string = args[0].String()
	return this.Call("getAttribute", selector)
})
