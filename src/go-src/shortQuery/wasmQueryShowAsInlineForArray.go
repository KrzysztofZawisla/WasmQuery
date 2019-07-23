package shortquery

import "syscall/js"

/* WasmQueryShowAsInlineBlock is a function from shortQuery module. This function set/return display inline-block for html elements. This function can only work on arrays of html elements. */
var WasmQueryShowAsInlineForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	var outputArray []interface{}
	for i := 0; i < this.Length(); i++ {
		this.Index(i).Get("style").Set("display", "inline")
		valueOfIteration := this.Index(i).Get("style").Get("display")
		outputArray = append(outputArray, valueOfIteration)
	}
	return outputArray
})
