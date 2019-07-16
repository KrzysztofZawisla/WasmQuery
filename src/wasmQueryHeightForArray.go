package main

import "syscall/js"

var wasmQueryHeightForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	var outputArray []interface{}
	if len(args) == 0 {
		for i := 0; i < this.Length(); i++ {
			valueOfIteration := this.Index(i).Get("style").Get("height")
			outputArray = append(outputArray, valueOfIteration)
		}
		return outputArray
	}
	value := args[0]
	if value.Type().String() == "object" {
		for i := 0; i < this.Length(); i++ {
			this.Index(i).Get("style").Set("height", value.Index(i))
			valueOfIteration := this.Index(i).Get("style").Get("height")
			outputArray = append(outputArray, valueOfIteration)
		}
	} else {
		for i := 0; i < this.Length(); i++ {
			this.Index(i).Get("style").Set("height", value)
			valueOfIteration := this.Index(i).Get("style").Get("height")
			outputArray = append(outputArray, valueOfIteration)
		}
	}
	return outputArray
})
