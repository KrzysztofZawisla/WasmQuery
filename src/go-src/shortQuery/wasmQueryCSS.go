package shortquery

import "syscall/js"

/* WasmQueryCSS is a function from shortQuery module. This function set/return css attribute for html element. This function can only work on single element. */
var WasmQueryCSS js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return js.Global().Get("window").Call("getComputedStyle", this)
	}
	var selector string = args[0].String()
	if len(args) == 1 {
		return js.Global().Get("window").Call("getComputedStyle", this).Get(selector)
	}
	var value string = args[1].String()
	this.Get("style").Set(selector, value)
	return js.Global().Get("window").Call("getComputedStyle", this).Get(selector)
})
