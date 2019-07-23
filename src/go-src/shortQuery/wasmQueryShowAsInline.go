package shortquery

import "syscall/js"

/* WasmQueryShowAsInline is a function from shortQuery module. This function set/return display inline for html element. This function can only work on single element. */
var WasmQueryShowAsInline js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	this.Get("style").Set("display", "inline")
	return this.Get("style").Get("display")
})
