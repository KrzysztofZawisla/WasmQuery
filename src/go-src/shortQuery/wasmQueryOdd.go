package shortquery

import (
	"syscall/js"
)

// WasmQueryOdd is a function from shortQuery module. This function return odd html elements. This function can only work on arrays of html elements.
func WasmQueryOdd(this js.Value, args []js.Value) interface{} {
	var outputArray []interface{}
	for i := 0; i < this.Length(); i++ {
		if i%2 != 0 {
			outputArray = append(outputArray, this.Index(i))
		}
	}
	if len(outputArray) == 0 {
		return outputArray
	}
	outputArrayJS := js.ValueOf(outputArray)
	// dodać inne funkcje
	outputArrayJS.Set("val", js.FuncOf(WasmQueryValForArray))
	return outputArrayJS
}
