package shortquery

import (
	"syscall/js"
)

// WasmQueryCSS is a function from shortQuery module. This function set/return css attribute for html element. This function can only work on single element.
func WasmQueryCSS(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		selector := args[0]
		if len(args) == 1 {
			if selector.Type() == js.TypeObject {
				var outputArray []interface{}
				for i := 0; i < selector.Length(); i++ {
					outputArray = append(outputArray, js.Global().Get("window").Call("getComputedStyle", this).Get(selector.Index(i).String()))
				}
				return outputArray
			}
			return js.Global().Get("window").Call("getComputedStyle", this).Get(selector.String())
		}
		var outputArray []interface{}
		value := args[1]
		if selector.Type() == js.TypeObject {
			if value.Type() == js.TypeObject {
				for i := 0; i < selector.Length(); i++ {
					for j := 0; j < value.Length(); j++ {
						if j == i {
							this.Get("style").Set(selector.Index(i).String(), value.Index(i).String())
							outputArray = append(outputArray, js.Global().Get("window").Call("getComputedStyle", this).Get(selector.Index(i).String()))
						}
					}
				}
			} else {
				for i := 0; i < selector.Length(); i++ {
					this.Get("style").Set(selector.Index(i).String(), value.String())
					outputArray = append(outputArray, js.Global().Get("window").Call("getComputedStyle", this).Get(selector.Index(i).String()))
				}
			}
			return outputArray
		}
		this.Get("style").Set(selector.String(), value.String())
		return js.Global().Get("window").Call("getComputedStyle", this).Get(selector.String())
	}
	return js.Global().Get("window").Call("getComputedStyle", this)
}
