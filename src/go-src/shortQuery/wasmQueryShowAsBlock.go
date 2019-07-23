package shortquery

import "syscall/js"

/* WasmQueryShow is a function from shortQuery module. This function set/return display block for html element. This function can only work on single element. */
var WasmQueryShowAsBlock js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	this.Get("style").Set("display", "block")
	return this.Get("style").Get("display")
})
