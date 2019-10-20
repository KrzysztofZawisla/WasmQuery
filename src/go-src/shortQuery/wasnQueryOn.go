package shortquery

import (
	"syscall/js"
)

// WasmQueryOn func
func WasmQueryOn(this js.Value, args []js.Value) interface{} {
	if len(args) > 1 {
		selector := args[0]
		function := args[1]
		if function.Type() == js.TypeFunction && selector.Type() == js.TypeString {
			this.Call("addEventListener", selector, function)
		} else {
			if function.Type() != js.TypeFunction && selector.Type() != js.TypeString {
				return js.Global().Get("Error").New("Wrongly passed arguments. Selector and function have the wrong type.")
			} else if function.Type() != js.TypeFunction {
				return js.Global().Get("Error").New("Wrongly passed arguments. Function have the wrong type.")
			} else if selector.Type() != js.TypeString {
				return js.Global().Get("Error").New("Wrongly passed arguments. Selector have the wrong type.")
			}
		}
	}
	return js.Null()
}
