package shortquery

import (
	"syscall/js"
)

var WasmQueryAddClass js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return this.Get("classList")
	}
	value := args[0]
	if value.Type().String() == "object" {
		for i := 0; i < value.Length(); i++ {
			this.Get("classList").Call("add", value.Index(i).String())
		}
	} else {
		this.Get("classList").Call("add", value.String())
	}
	return this.Get("classList")
})
