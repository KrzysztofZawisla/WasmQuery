package shortquery

import "syscall/js"

// WasmQueryVal is a function from shortQuery module. This function set/return value attribute for html element. This function can only work on single element.
func WasmQueryVal(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0]
		this.Set("value", value)
	}
	return this.Get("value")
}
