package main

import "syscall/js"

var wasmQueryAttrForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return nil
	}
	var selector string = args[0].String()
	var outputArray []interface{}
	if len(args) >= 1 {
		for i := 0; i < this.Length(); i++ {
			valueOfIteration := this.Index(i).Call("getAttribute", selector)
			outputArray = append(outputArray, valueOfIteration)
		}
		return outputArray
	}
	return nil
})
