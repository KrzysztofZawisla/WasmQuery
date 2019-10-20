package shortquery

import (
	"syscall/js"
)

// WasmQueryAttrForArray is a function from shortQuery module. This function set/return class/classes for html elements. This function can only work on arrays of html elements.
func WasmQueryAttrForArray(this js.Value, args []js.Value) interface{} {
	var outputArray []interface{}
	jsOutputArray := js.ValueOf(outputArray)
	if len(args) > 0 {
		selector := args[0]
		if len(args) == 1 {
			for i := 0; i < this.Length(); i++ {
				jsValueArgs := []js.Value{selector}
				jsOutputArray.SetIndex(i, WasmQueryAttr(this.Index(i), jsValueArgs))
			}
		} else {
			value := args[1]
			if selector.Type() == js.TypeObject {
				for i := 0; i < this.Length(); i++ {
					jsValueArgs := []js.Value{selector, value}
					jsOutputArray.SetIndex(i, WasmQueryAttr(this.Index(i), jsValueArgs))
				}
			} else {
				if value.Type() == js.TypeObject {
					for i := 0; i < value.Length(); i++ {
						for j := 0; j < this.Length(); j++ {
							if j == i {
								jsValueArgs := []js.Value{selector, value.Index(i)}
								jsOutputArray.SetIndex(i, WasmQueryAttr(this.Index(i), jsValueArgs))
							}
						}
					}
				} else {
					for i := 0; i < this.Length(); i++ {
						jsValueArgs := []js.Value{selector, value}
						jsOutputArray.SetIndex(i, WasmQueryAttr(this.Index(i), jsValueArgs))
					}
				}
			}
		}
	} else {
		for i := 0; i < this.Length(); i++ {
			jsOutputArray.SetIndex(i, WasmQueryAttr(this.Index(i), args))
		}
	}
	return jsOutputArray
}
