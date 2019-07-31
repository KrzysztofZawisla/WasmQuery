package shortquery

import (
	"syscall/js"
)

var WasmQueryHasClassForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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
			var arrayForOneValue []interface{}
			for j := 0; j < value.Length(); j++ {
				valueOfIteration := this.Index(i).Get("classList").Call("contains", value.Index(j))
				arrayForOneValue = append(arrayForOneValue, valueOfIteration)
			}
			outputArray = append(outputArray, arrayForOneValue)
		}
		return outputArray
	} else {
		for i := 0; i < this.Length(); i++ {
			var arrayForOneValue []interface{}
			for j := 0; j < value.Length(); j++ {
				valueOfIteration := this.Index(i).Get("classList").Call("contains", value)
				arrayForOneValue = append(arrayForOneValue, valueOfIteration)
			}
			outputArray = append(outputArray, arrayForOneValue)
		}
		return outputArray
	}
	return outputArray
})
