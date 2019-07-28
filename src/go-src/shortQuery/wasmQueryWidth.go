package shortquery

import "syscall/js"

/* WasmQueryWidth is a function from shortQuery module. This function set/return width attribute for html element. This function can only work on single element. */
var WasmQueryWidth js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return this.Get("style").Get("width")
	}
	value := args[0]
	this.Get("style").Set("width", value)
	return js.Global().Get("window").Call("getComputedStyle", this).Get("width")
})
