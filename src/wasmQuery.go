package main

import (
	"syscall/js"
)

var wasmQuery js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return nil
	}
	domSelector := args[0]
	if domSelector == js.Global().Get("document") {
		returnValue := args[0]
		return returnValue
	}
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
		returnValue.Index(i).Set("width", wasmQueryWidth)
		returnValue.Index(i).Set("height", wasmQueryHeight)
		returnValue.Index(i).Set("attr", wasmQueryAttr)
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
	returnValue.Set("width", wasmQueryWidthForArray)
	returnValue.Set("height", wasmQueryHeightForArray)
	returnValue.Set("attr", wasmQueryAttrForArray)
	return returnValue
})
