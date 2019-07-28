package shortquery

import "syscall/js"

/* WasmQueryCSSForArray is a function from shortQuery module. This function set/return css attribute for html elements. This function can only work on arrays of html elements. */
var WasmQueryCSSForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	var outputArray []interface{}
	if len(args) == 0 {
		for i := 0; i < this.Length(); i++ {
			valueOfIteration := js.Global().Get("window").Call("getComputedStyle", this.Index(i))
			outputArray = append(outputArray, valueOfIteration)
		}
		return outputArray
	}
	var selector string = args[0].String()
	if len(args) == 1 {
		for i := 0; i < this.Length(); i++ {
			valueOfIteration := js.Global().Get("window").Call("getComputedStyle", this.Index(i)).Get(selector)
			outputArray = append(outputArray, valueOfIteration)
		}
		return outputArray
	}
	value := args[1]
	if value.Type().String() == "object" {
		for i := 0; i < this.Length(); i++ {
			this.Index(i).Get("style").Set(selector, value.Index(i).String())
			valueOfIteration := js.Global().Get("window").Call("getComputedStyle", this.Index(i)).Get(selector)
			outputArray = append(outputArray, valueOfIteration)
		}
	} else {
		for i := 0; i < this.Length(); i++ {
			this.Index(i).Get("style").Set(selector, value.String())
			valueOfIteration := js.Global().Get("window").Call("getComputedStyle", this.Index(i)).Get(selector)
			outputArray = append(outputArray, valueOfIteration)
		}
	}
	return outputArray
})
