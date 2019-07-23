package shortquery

import "syscall/js"

/* WasmQueryShowAsFlex is a function from shortQuery module. This function set/return display flex for html element. This function can only work on single element. */
var WasmQueryShowAsFlex js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	this.Get("style").Set("display", "flex")
	return this.Get("style").Get("display")
})
