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
			for j := 0; j < value.Length(); j++ {
				this.Index(i).Get("classList").Call("add", value.Index(j))
			}
			valueOfIteration := this.Index(i).Get("classList")
			outputArray = append(outputArray, valueOfIteration)
		}
	} else {
		for i := 0; i < this.Length(); i++ {
			this.Index(i).Get("classList").Call("add", value)
			valueOfIteration := this.Index(i).Get("classList")
			outputArray = append(outputArray, valueOfIteration)
		}
	}
	return outputArray
})
