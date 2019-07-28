package shortquery

import (
	"syscall/js"
)

var WasmQueryAddClassForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	var outputArray []interface{}
	if len(args) == 0 {
		for i := 0; i < this.Length(); i++ {
			valueOfIteration := this.Index(i).Get("classList")
			outputArray = append(outputArray, valueOfIteration)
		}
		return outputArray
	}
	value := args[0]
	if value.Type().String() == "object" {
		for i := 0; i < this.Length(); i++ {
			if value.Index(i).String() == "undefined" || value.Index(i).String() == "" {
				continue
			}
			this.Index(i).Get("classList").Call("add", value.Index(i))
			valueOfIteration := this.Index(i).Get("classList")
			outputArray = append(outputArray, valueOfIteration)
		}
	} else {
		for i := 0; i < this.Length(); i++ {
			if value.String() == "undefined" || value.String() == "" {
				continue
			}
			this.Index(i).Get("classList").Call("add", value)
			valueOfIteration := this.Index(i).Get("classList")
			outputArray = append(outputArray, valueOfIteration)
		}
	}
	return outputArray
})
