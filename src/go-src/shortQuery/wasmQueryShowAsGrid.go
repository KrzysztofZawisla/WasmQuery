package shortquery

import "syscall/js"

// WasmQueryShowAsGrid is a function from shortQuery module. This function set/return display block for html element. This function can only work on single element.
func WasmQueryShowAsGrid(this js.Value, args []js.Value) interface{} {
	this.Get("style").Set("display", "grid")
	return this.Get("style").Get("display")
}
