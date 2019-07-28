package shortquery

import "syscall/js"

/* WasmQueryHeightForArray is a function from shortQuery module. This function set display none and return display none value for html elements. This function can only work on arrays of html elements. */
var WasmQueryHideForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	var outputArray []interface{}
	for i := 0; i < this.Length(); i++ {
		this.Index(i).Get("style").Set("display", "none")
		valueOfIteration := js.Global().Get("window").Call("getComputedStyle", this.Index(i)).Get("display")
		outputArray = append(outputArray, valueOfIteration)
	}
	return outputArray
})
