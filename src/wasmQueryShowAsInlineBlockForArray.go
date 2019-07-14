package main

import "syscall/js"

var wasmQueryShowAsInlineBlockForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	var outputArray []interface{}
	for i := 0; i < this.Length(); i++ {
		this.Index(i).Get("style").Set("display", "inline-block")
		valueOfIteration := this.Index(i).Get("style").Get("display")
		outputArray = append(outputArray, valueOfIteration)
	}
	return outputArray
})
