package shortquery

import (
	"syscall/js"
)

// WasmQueryToggleClassForArray is a function from shortQuery module. This function set/return value attribute for html elements. This function can only work on arrays of html elements.
func WasmQueryToggleClassForArray(this js.Value, args []js.Value) interface{} {
	var outputArray []interface{}
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeObject {
			for i := 0; i < value.Length(); i++ {
				jsValueArgs := []js.Value{value.Index(i)}
				for j := 0; j < this.Length(); j++ {
					if j == i {
						outputArray = append(outputArray, WasmQueryToggleClass(this.Index(i), jsValueArgs))
					}
				}

			}
		} else {
			for i := 0; i < this.Length(); i++ {
				jsValueArgs := []js.Value{value}
				outputArray = append(outputArray, WasmQueryToggleClass(this.Index(i), jsValueArgs))
			}
		}
	} else {
		for i := 0; i < this.Length(); i++ {
			outputArray = append(outputArray, WasmQueryToggleClass(this.Index(i), args))
		}
	}
	return outputArray
}
