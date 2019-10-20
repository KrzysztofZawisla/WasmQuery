package shortquery

import "syscall/js"

// WasmQueryShowAsFlex is a function from shortQuery module. This function set/return display block for html element. This function can only work on single element.
func WasmQueryShowAsFlex(this js.Value, args []js.Value) interface{} {
	this.Get("style").Set("display", "flex")
	return this.Get("style").Get("display")
}
