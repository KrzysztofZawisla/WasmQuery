package shortquery

import "syscall/js"

/* WasmQueryShowAsInline is a function from shortQuery module. This function set/return display inline-block for html element. This function can only work on single element. */
var WasmQueryShowAsInlineBlock js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	this.Get("style").Set("display", "inline-block")
	return js.Global().Get("window").Call("getComputedStyle", this).Get("display")
})
