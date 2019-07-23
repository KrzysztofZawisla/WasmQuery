package shortquery

import "syscall/js"

/* WasmQueryCSSForArray is a function from shortQuery module. This function set/return css attribute for html elements. This function can only work on arrays of html elements. */
var WasmQueryCSSForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return nil
	}
	var selector string = args[0].String()
	var outputArray []interface{}
	if len(args) == 1 {
		for i := 0; i < this.Length(); i++ {
			valueOfIteration := this.Index(i).Get("style").Get(selector)
			outputArray = append(outputArray, valueOfIteration)
		}
		return outputArray
	} else if len(args) >= 2 {
		value := args[1]
		if value.Type().String() == "object" {
			for i := 0; i < this.Length(); i++ {
				this.Index(i).Get("style").Set(selector, value.Index(i).String())
				valueOfIteration := this.Index(i).Get("style").Get(selector)
				outputArray = append(outputArray, valueOfIteration)
			}
		} else {
			for i := 0; i < this.Length(); i++ {
				this.Index(i).Get("style").Set(selector, value.String())
				valueOfIteration := this.Index(i).Get("style").Get(selector)
				outputArray = append(outputArray, valueOfIteration)
			}
		}
		return outputArray
	}
	return nil
})
