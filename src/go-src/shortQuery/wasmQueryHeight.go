package shortquery

import "syscall/js"

// WasmQueryHeight is a function from shortQuery module. This function sets or returns height attribute of html element. Can work only on a single element. Never returns undefined.
func WasmQueryHeight(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeObject {
			for i := 0; i < value.Length(); i++ {
				if value.Type() == js.TypeNumber {
					this.Get("style").Set("height", value.Index(i).String()+"px")
				} else {
					this.Get("style").Set("height", value.Index(i))
				}
			}
		} else if value.Type() == js.TypeNumber {
			this.Get("style").Set("height", value.String()+"px")
		} else {
			this.Get("style").Set("height", value)
		}
	}
	return js.Global().Get("window").Call("getComputedStyle", this).Get("height")
}
