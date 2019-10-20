package shortquery

import (
	"syscall/js"
)

// WasmQueryAttr is a function from shortQuery module. This function set/return attribute for html element. This function can only work on single element.
func WasmQueryAttr(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		selector := args[0]
		if len(args) == 1 {
			if selector.Type() == js.TypeObject {
				var outputArray []interface{}
				jsOutputArray := js.ValueOf(outputArray)
				for i := 0; i < selector.Length(); i++ {
					jsOutputArray.SetIndex(i, this.Call("getAttribute", selector.Index(i)))
				}
				return jsOutputArray
			}
			return this.Call("getAttribute", selector)
		}
		var outputArray []interface{}
		jsOutputArray := js.ValueOf(outputArray)
		value := args[1]
		if selector.Type() == js.TypeObject {
			if value.Type() == js.TypeObject {
				for i := 0; i < selector.Length(); i++ {
					for j := 0; j < value.Length(); j++ {
						if j == i {
							this.Call("setAttribute", selector.Index(i), value.Index(i))
							jsOutputArray.SetIndex(i, this.Call("getAttribute", selector.Index(i)))
						}
					}
				}
			} else {
				for i := 0; i < selector.Length(); i++ {
					this.Call("setAttribute", selector.Index(i), value)
					jsOutputArray.SetIndex(i, this.Call("getAttribute", selector.Index(i)))
				}
			}
			return jsOutputArray
		}
		this.Call("setAttribute", selector, value)
		return this.Call("getAttribute", selector)
	}
	return this.Get("attributes")
}
