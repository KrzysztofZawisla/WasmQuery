package shortquery

import (
	"syscall/js"
)

/* WasmQueryAttrForArray is a function from shortQuery module. This function set/return attribute for html elements. This function can only work on arrays of html elements. */
var WasmQueryAttrForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return nil
	}
	selector := args[0]
	if len(args) >= 1 {
		if selector.Type().String() == "object" {
			var outputArray []interface{}
			for i := 0; i < this.Length(); i++ {
				var arrayForOneValue []interface{}
				for j := 0; j < selector.Length(); j++ {
					valueOfIteration := this.Index(i).Call("getAttribute", selector.Index(j))
					arrayForOneValue = append(arrayForOneValue, valueOfIteration)
				}
				outputArray = append(outputArray, arrayForOneValue)
			}
			return outputArray
		}
		var outputArray []interface{}
		for i := 0; i < this.Length(); i++ {
			valueOfIteration := this.Index(i).Call("getAttribute", selector.String())
			outputArray = append(outputArray, valueOfIteration)
		}
		return outputArray
	}
	return nil
})
