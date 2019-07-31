package shortquery

import (
	"syscall/js"
)

var WasmQueryToggleClassForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	var outputArray []interface{}
	if len(args) == 0 {
		for i := 0; i < this.Length(); i++ {
			valueOfIteration := this.Index(i).Get("classList")
			outputArray = append(outputArray, valueOfIteration)
		}
		return outputArray
	}
	selector := args[0]
	if selector.Type().String() == "object" {
		for i := 0; i < this.Length(); i++ {
			for j := 0; j < selector.Length(); j++ {
				var checkerOutput bool = this.Index(i).Get("classList").Call("contains", selector.Index(j)).Bool()
				if checkerOutput == true {
					this.Index(i).Get("classList").Call("remove", selector.Index(j))
				} else {
					this.Index(i).Get("classList").Call("add", selector.Index(j))
				}
			}
			valueOfIteration := this.Index(i).Get("classList")
			outputArray = append(outputArray, valueOfIteration)
		}
	} else {
		for i := 0; i < this.Length(); i++ {
			if selector.String() == "undefined" || selector.String() == "" {
				continue
			}
			var checkerOutput bool = this.Index(i).Get("classList").Call("contains", selector).Bool()
			if checkerOutput == true {
				this.Index(i).Get("classList").Call("remove", selector)
			} else {
				this.Index(i).Get("classList").Call("add", selector)
			}
			valueOfIteration := this.Index(i).Get("classList")
			outputArray = append(outputArray, valueOfIteration)
		}
	}
	return outputArray
})
