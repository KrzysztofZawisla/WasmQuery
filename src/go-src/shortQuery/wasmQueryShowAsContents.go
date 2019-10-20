package shortquery

import "syscall/js"

// WasmQueryShowAsContents is a function from shortQuery module. This function set/return display block for html element. This function can only work on single element.
func WasmQueryShowAsContents(this js.Value, args []js.Value) interface{} {
	this.Get("style").Set("display", "contents")
	return this.Get("style").Get("display")
}
