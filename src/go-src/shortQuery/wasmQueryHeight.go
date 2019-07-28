package shortquery

import "syscall/js"

/* WasmQueryHeight is a function from shortQuery module. This function set/return height for html element. This function can only work on single element. */
var WasmQueryHeight js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return js.Global().Get("window").Call("getComputedStyle", this).Get("height")
	}
	value := args[0]
	this.Get("style").Set("height", value)
	return js.Global().Get("window").Call("getComputedStyle", this).Get("height")
})
