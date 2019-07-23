package shortquery

import (
	"syscall/js"
)

/* WasmQuery is a constructor function of shortQuery module. */
var WasmQuery js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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
		returnValue.Index(i).Set("css", WasmQueryCSS)
		returnValue.Index(i).Set("hide", WasmQueryHide)
		returnValue.Index(i).Set("show", WasmQueryShow)
		returnValue.Index(i).Set("showAsBlock", WasmQueryShowAsBlock)
		returnValue.Index(i).Set("showAsInline", WasmQueryShowAsInline)
		returnValue.Index(i).Set("showAsInlineBlock", WasmQueryShowAsInlineBlock)
		returnValue.Index(i).Set("showAsFlex", WasmQueryShowAsFlex)
		returnValue.Index(i).Set("val", WasmQueryVal)
		returnValue.Index(i).Set("width", WasmQueryWidth)
		returnValue.Index(i).Set("height", WasmQueryHeight)
		returnValue.Index(i).Set("attr", WasmQueryAttr)
	}
	returnValue.Set("css", WasmQueryCSSForArray)
	returnValue.Set("hide", WasmQueryHideForArray)
	returnValue.Set("show", WasmQueryShowForArray)
	returnValue.Set("showAsBlock", WasmQueryShowAsBlockForArray)
	returnValue.Set("showAsInline", WasmQueryShowAsInlineForArray)
	returnValue.Set("showAsInlineBlock", WasmQueryShowAsInlineBlockForArray)
	returnValue.Set("showAsFlex", WasmQueryShowAsFlexForArray)
	returnValue.Set("val", WasmQueryValForArray)
	returnValue.Set("len", WasmQueryLen)
	returnValue.Set("width", WasmQueryWidthForArray)
	returnValue.Set("height", WasmQueryHeightForArray)
	returnValue.Set("attr", WasmQueryAttrForArray)
	return returnValue
})
