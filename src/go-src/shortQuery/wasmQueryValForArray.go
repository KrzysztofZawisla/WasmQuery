package shortquery

import (
	"syscall/js"
)

/* WasmQueryValForArray is a function from shortQuery module. This function set/return value attribute for html elements. This function can only work on arrays of html elements. */
var WasmQueryValForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	var outputArray []interface{}
	if len(args) == 0 {
		for i := 0; i < this.Length(); i++ {
			valueOfIteration := this.Index(i).Get("value")
			outputArray = append(outputArray, valueOfIteration)
		}
		return outputArray
	}
	value := args[0]
	if value.Type().String() == "object" {
		for i := 0; i < this.Length(); i++ {
			this.Index(i).Set("value", value.Index(i))
			valueOfIteration := this.Index(i).Get("value")
			outputArray = append(outputArray, valueOfIteration)
		}
	} else {
		for i := 0; i < this.Length(); i++ {
			this.Index(i).Set("value", value)
			valueOfIteration := this.Index(i).Get("value")
			outputArray = append(outputArray, valueOfIteration)
		}
	}
	return outputArray
})
