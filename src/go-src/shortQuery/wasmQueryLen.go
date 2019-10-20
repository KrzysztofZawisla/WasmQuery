package shortquery

import "syscall/js"

// WasmQueryLen is a function from shortQuery module. This function returns the length of html elements array. Can only works on an array of html element.
func WasmQueryLen(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(this.Get("length"))
}
