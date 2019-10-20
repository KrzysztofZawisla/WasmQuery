package shortquery

import (
	"syscall/js"
)

// WasmQueryRemoveAttr is a function from shortQuery module. This function set/return class/classes for html element. This function can only work on single element.
func WasmQueryRemoveAttr(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeObject {
			for i := 0; i < value.Length(); i++ {
				this.Call("removeAttribute", value.Index(i))
			}
		}
		this.Call("removeAttribute", value)
	}
	return this.Get("attributes")
}
