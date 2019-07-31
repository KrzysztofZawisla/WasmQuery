package shortquery

import (
	"syscall/js"
)

/* WasmQueryAttrForArray is a function from shortQuery module. This function set/return attribute for html elements. This function can only work on arrays of html elements. */
var WasmQueryAttrForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return nil
	}
	value := args[0]
	var outputArray []interface{}
	if value.Type().String() == "object" {
		for i := 0; i < this.Length(); i++ {
			var arrayForOneValue []interface{}
			for j := 0; j < value.Length(); j++ {
				valueOfIteration := this.Get("classList").Call("contains", value)
				arrayForOneValue = append(arrayForOneValue, valueOfIteration)
			}
			outputArray = append(outputArray, arrayForOneValue)
		}
		return outputArray
	}
	for i := 0; i < this.Length(); i++ {
		valueOfIteration := this.Index(i).Call("getAttribute", value.String())
		outputArray = append(outputArray, valueOfIteration)
	}
	return outputArray
})
