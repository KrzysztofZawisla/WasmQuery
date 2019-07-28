package shortquery

import (
	"syscall/js"
)

var WasmQueryRemoveClassForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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
			this.Index(i).Get("classList").Call("remove", value.Index(i).String())
			valueOfIteration := this.Index(i).Get("classList")
			outputArray = append(outputArray, valueOfIteration)
		}
	} else {
		for i := 0; i < this.Length(); i++ {
			this.Index(i).Get("classList").Call("remove", value.String())
			valueOfIteration := this.Index(i).Get("classList")
			outputArray = append(outputArray, valueOfIteration)
		}
	}
	return outputArray
})
