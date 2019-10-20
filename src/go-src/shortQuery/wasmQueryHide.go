package shortquery

import "syscall/js"

// WasmQueryHide is a function from shortQuery module. This function set display none and return display none value for html element. This function can only work on single element.
func WasmQueryHide(this js.Value, args []js.Value) interface{} {
	this.Get("style").Set("display", "none")
	return this.Get("style").Get("display")
}
