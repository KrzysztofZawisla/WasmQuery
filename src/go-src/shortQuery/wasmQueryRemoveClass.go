package shortquery

import (
	"syscall/js"
)

// WasmQueryRemoveClass is a function from shortQuery module. This function remove/return class/classes for html element. This function can only work on single element.
func WasmQueryRemoveClass(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeObject {
			for i := 0; i < value.Length(); i++ {
				this.Get("classList").Call("remove", value.Index(i))
			}
		} else {
			this.Get("classList").Call("remove", value)
		}
	}
	return this.Get("classList")
}
