package shortquery

import "syscall/js"

// WasmQueryShowAsInlineGridForArray is a function from shortQuery module. This function set display none and return display none value for html elements. This function can only work on arrays of html elements.
func WasmQueryShowAsInlineGridForArray(this js.Value, args []js.Value) interface{} {
	var outputArray []interface{}
	for i := 0; i < this.Length(); i++ {
		outputArray = append(outputArray, WasmQueryShowAsInlineGrid(this.Index(i), args))
	}
	return outputArray
}
