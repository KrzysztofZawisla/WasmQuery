package shortquery

import (
	"syscall/js"
)

// WasmQueryHrefForArray is a function from shortQuery module. This function set/return class/classes for html elements. This function can only work on arrays of html elements.
func WasmQueryHrefForArray(this js.Value, args []js.Value) interface{} {
	var outputArray []interface{}
	jsOutputArray := js.ValueOf(outputArray)
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeObject {
			for i := 0; i < value.Length(); i++ {
				jsValueArgs := []js.Value{value.Index(i)}
				for j := 0; j < this.Length(); j++ {
					if j == i {
						jsOutputArray.SetIndex(i, WasmQueryHref(this.Index(i), jsValueArgs))
					}
				}
			}
		} else {
			for i := 0; i < this.Length(); i++ {
				jsValueArgs := []js.Value{value}
				jsOutputArray.SetIndex(i, WasmQueryHref(this.Index(i), jsValueArgs))
			}
		}
	} else {
		for i := 0; i < this.Length(); i++ {
			jsOutputArray.SetIndex(i, WasmQueryHref(this.Index(i), args))
		}
	}
	return jsOutputArray
}
