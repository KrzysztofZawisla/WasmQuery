package shortquery

import "syscall/js"

// WasmQueryWidth is a function from shortQuery module. This function sets or returns width attribute of html element. Can work only on a single element. Never returns undefined.
func WasmQueryWidth(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeObject {
			for i := 0; i < value.Length(); i++ {
				if value.Type() == js.TypeNumber {
					this.Get("style").Set("width", value.Index(i).String()+"px")
				} else {
					this.Get("style").Set("width", value.Index(i))
				}
			}
		} else if value.Type() == js.TypeNumber {
			this.Get("style").Set("width", value.String()+"px")
		} else {
			this.Get("style").Set("width", value)
		}
	}
	return js.Global().Get("window").Call("getComputedStyle", this).Get("width")
}
