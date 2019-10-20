package shortquery

import (
	"syscall/js"
)

// WasmQueryHasClass is a function from shortQuery module. This function set/return class/classes for html element. This function can only work on single element.
func WasmQueryHasClass(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeObject {
			var outputArray []interface{}
			for i := 0; i < value.Length(); i++ {
				outputArray = append(outputArray, this.Get("classList").Call("contains", value.Index(i)))
			}
			return outputArray
		}
		return this.Get("classList").Call("contains", value)
	}
	return this.Get("classList")
}
