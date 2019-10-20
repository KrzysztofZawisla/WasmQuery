package shortquery

import "syscall/js"

// WasmQueryShowAsInitial is a function from shortQuery module. This function set/return display block/inline by smart checking for html element. This function can only work on single element.
func WasmQueryShowAsInitial(this js.Value, args []js.Value) interface{} {
	this.Get("style").Set("display", "initial")
	return this.Get("style").Get("display")
}
