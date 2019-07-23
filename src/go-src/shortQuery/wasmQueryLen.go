package shortquery

import "syscall/js"

/* WasmQueryLen is a function from shortQuery module. This function return length of html elements array. This function can only work on arrays of html elements. */
var WasmQueryLen js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	return this.Get("length")
})
