package shortquery

import "syscall/js"

/* WasmQueryShowForArray is a function from shortQuery module. This function set/return display block for html elements. This function can only work on arrays of html elements. */
var WasmQueryShowAsBlockForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	var outputArray []interface{}
	for i := 0; i < this.Length(); i++ {
		this.Index(i).Get("style").Set("display", "block")
		valueOfIteration := this.Index(i).Get("style").Get("display")
		outputArray = append(outputArray, valueOfIteration)
	}
	return outputArray
})
