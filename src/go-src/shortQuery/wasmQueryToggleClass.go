package shortquery

import (
	"syscall/js"
)

var WasmQueryToggleClass js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return nil
	}
	var selector string = args[0].String()
	var checkerOutput bool = this.Get("classList").Call("contains", selector).Bool()
	if checkerOutput == true {
		this.Get("classList").Call("remove", selector)
	} else {
		this.Get("classList").Call("add", selector)
	}
	return this.Get("classList")
})
