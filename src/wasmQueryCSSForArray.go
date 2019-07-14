package main

import "syscall/js"

var wasmQueryCSSForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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
		var value string = args[1].String()
		for i := 0; i < this.Length(); i++ {
			this.Index(i).Get("style").Set(selector, value)
			valueOfIteration := this.Index(i).Get("style").Get(selector)
			outputArray = append(outputArray, valueOfIteration)
		}
		return outputArray
	}
	return nil
})
