package shortquery

import "syscall/js"

// WasmQueryShowAsInlineBlock is a function from shortQuery module. This function set/return display block for html element. This function can only work on single element.
func WasmQueryShowAsInlineBlock(this js.Value, args []js.Value) interface{} {
	this.Get("style").Set("display", "inline-block")
	return this.Get("style").Get("display")
}
