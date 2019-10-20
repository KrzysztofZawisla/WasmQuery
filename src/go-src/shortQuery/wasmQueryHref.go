package shortquery

import (
	"syscall/js"
)

// WasmQueryHref is a function from shortQuery module. This function set/return innerHTML for html element. This function can only work on single element.
func WasmQueryHref(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0]
		this.Set("href", value)
	}
	return this.Get("href")
}
