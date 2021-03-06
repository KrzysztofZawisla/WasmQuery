package shortquery

import "syscall/js"

// WasmQueryShowAsListItem is a function from shortQuery module. This function set/return display block for html element. This function can only work on single element.
func WasmQueryShowAsListItem(this js.Value, args []js.Value) interface{} {
	this.Get("style").Set("display", "list-item")
	return this.Get("style").Get("display")
}
