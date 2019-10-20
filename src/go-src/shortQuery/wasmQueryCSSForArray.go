package shortquery

import (
	"syscall/js"
)

// WasmQueryCSSForArray is a function from shortQuery module. This function set/return class/classes for html elements. This function can only work on arrays of html elements.
func WasmQueryCSSForArray(this js.Value, args []js.Value) interface{} {
	var outputArray []interface{}
	if len(args) > 0 {
		selector := args[0]
		if len(args) == 1 {
			for i := 0; i < this.Length(); i++ {
				jsValueArgs := []js.Value{selector}
				outputArray = append(outputArray, WasmQueryCSS(this.Index(i), jsValueArgs))
			}
		} else {
			value := args[1]
			if selector.Type() == js.TypeObject {
				for i := 0; i < this.Length(); i++ {
					jsValueArgs := []js.Value{selector, value}
					outputArray = append(outputArray, WasmQueryCSS(this.Index(i), jsValueArgs))
				}
			} else {
				if value.Type() == js.TypeObject {
					for i := 0; i < value.Length(); i++ {
						for j := 0; j < this.Length(); j++ {
							if j == i {
								jsValueArgs := []js.Value{selector, value.Index(i)}
								outputArray = append(outputArray, WasmQueryCSS(this.Index(i), jsValueArgs))
							}
						}
					}
				} else {
					for i := 0; i < this.Length(); i++ {
						jsValueArgs := []js.Value{selector, value}
						outputArray = append(outputArray, WasmQueryCSS(this.Index(i), jsValueArgs))
					}
				}
			}
		}
	} else {
		for i := 0; i < this.Length(); i++ {
			outputArray = append(outputArray, WasmQueryCSS(this.Index(i), args))
		}
	}
	return outputArray
}
