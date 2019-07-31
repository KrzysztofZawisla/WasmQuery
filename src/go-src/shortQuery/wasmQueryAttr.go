package shortquery

import (
	"syscall/js"
)

/* WasmQueryAttr is a function from shortQuery module. This function set/return attribute for html element. This function can only work on single element. */
var WasmQueryAttr js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return nil
	}
	value := args[0]
	if value.Type().String() == "object" {
		var outputArray []interface{}
		for i := 0; i < value.Length(); i++ {
			valueOfIteration := this.Call("getAttribute", value.Index(i).String())
			outputArray = append(outputArray, valueOfIteration)
		}
		return outputArray
	}
	return this.Call("getAttribute", value.String())
})
