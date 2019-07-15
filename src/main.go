package main

import (
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("$", wasmQuery)
	js.Global().Get("$").Set("disableLibrary", disableLibrary)
	<-c
}

var wasmQuery js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return nil
	}
	domSelector := args[0].String()
	returnValue := js.Global().Get("document").Call("querySelectorAll", domSelector)
	for i := 0; i < returnValue.Length(); i++ {
		returnValue.Index(i).Set("css", wasmQueryCSS)
		returnValue.Index(i).Set("hide", wasmQueryHide)
		returnValue.Index(i).Set("show", wasmQueryShow)
		returnValue.Index(i).Set("showAsBlock", wasmQueryShowAsBlock)
		returnValue.Index(i).Set("showAsInline", wasmQueryShowAsInline)
		returnValue.Index(i).Set("showAsInlineBlock", wasmQueryShowAsInlineBlock)
		returnValue.Index(i).Set("showAsFlex", wasmQueryShowAsFlex)
		returnValue.Index(i).Set("val", wasmQueryVal)
	}
	returnValue.Set("css", wasmQueryCSSForArray)
	returnValue.Set("hide", wasmQueryHideForArray)
	returnValue.Set("show", wasmQueryShowForArray)
	returnValue.Set("showAsBlock", wasmQueryShowAsBlockForArray)
	returnValue.Set("showAsInline", wasmQueryShowAsInlineForArray)
	returnValue.Set("showAsInlineBlock", wasmQueryShowAsInlineBlockForArray)
	returnValue.Set("showAsFlex", wasmQueryShowAsFlexForArray)
	returnValue.Set("val", wasmQueryValForArray)
	returnValue.Set("len", wasmQueryLen)
	return returnValue
})
