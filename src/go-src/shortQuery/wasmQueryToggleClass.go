package shortquery

import (
	"syscall/js"
)

// WasmQueryToggleClass is a function from shortQuery module. This function toggle/return class/classes for html element. This function can only work on single element.
func WasmQueryToggleClass(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeObject {
			for i := 0; i < value.Length(); i++ {
				checkerOutput := this.Get("classList").Call("contains", value.Index(i)).Bool()
				if checkerOutput == true {
					this.Get("classList").Call("remove", value.Index(i))
				} else {
					this.Get("classList").Call("add", value.Index(i))
				}
			}
		} else {
			checkerOutput := this.Get("classList").Call("contains", value).Bool()
			if checkerOutput == true {
				this.Get("classList").Call("remove", value)
			} else {
				this.Get("classList").Call("add", value)
			}
		}

	}
	return this.Get("classList")
}
