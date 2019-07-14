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
	var domSelector string = args[0].String()
	if string(domSelector[0]) == "#" || string(domSelector[0]) == "." {
		var selectorWithoutSymbol string = ""
		for pos, char := range domSelector {
			if pos == 0 {
				continue
			}
			selectorWithoutSymbol = selectorWithoutSymbol + string(char)
		}
		if string(domSelector[0]) == "#" {
			returnValue := js.Global().Get("document").Call("getElementById", selectorWithoutSymbol)
			returnValue.Set("css", wasmQueryCSS)
			returnValue.Set("hide", wasmQueryHide)
			returnValue.Set("show", wasmQueryShow)
			returnValue.Set("showAsBlock", wasmQueryShowAsBlock)
			returnValue.Set("showAsInline", wasmQueryShowAsInline)
			returnValue.Set("showAsInlineBlock", wasmQueryShowAsInlineBlock)
			returnValue.Set("showAsFlex", wasmQueryShowAsFlex)
			return returnValue
		} else if string(domSelector[0]) == "." {
			returnValue := js.Global().Get("document").Call("getElementsByClassName", selectorWithoutSymbol)
			for i := 0; i < returnValue.Length(); i++ {
				returnValue.Index(i).Set("css", wasmQueryCSS)
				returnValue.Index(i).Set("hide", wasmQueryHide)
				returnValue.Index(i).Set("show", wasmQueryShow)
				returnValue.Index(i).Set("showAsBlock", wasmQueryShowAsBlock)
				returnValue.Index(i).Set("showAsInline", wasmQueryShowAsInline)
				returnValue.Index(i).Set("showAsInlineBlock", wasmQueryShowAsInlineBlock)
				returnValue.Index(i).Set("showAsFlex", wasmQueryShowAsFlex)
			}
			returnValue.Set("css", wasmQueryCSSForArray)
			returnValue.Set("hide", wasmQueryHideForArray)
			returnValue.Set("show", wasmQueryShowForArray)
			returnValue.Set("showAsBlock", wasmQueryShowAsBlockForArray)
			returnValue.Set("showAsInline", wasmQueryShowAsInlineForArray)
			returnValue.Set("showAsInlineBlock", wasmQueryShowAsInlineBlockForArray)
			returnValue.Set("showAsFlex", wasmQueryShowAsFlexForArray)
			return returnValue
		}
	} else {
		returnValue := js.Global().Get("document").Call("getElementsByTagName", domSelector)
		for i := 0; i < returnValue.Length(); i++ {
			returnValue.Index(i).Set("css", wasmQueryCSS)
			returnValue.Index(i).Set("hide", wasmQueryHide)
			returnValue.Index(i).Set("show", wasmQueryShow)
			returnValue.Index(i).Set("showAsBlock", wasmQueryShowAsBlock)
			returnValue.Index(i).Set("showAsInline", wasmQueryShowAsInline)
			returnValue.Index(i).Set("showAsInlineBlock", wasmQueryShowAsInlineBlock)
			returnValue.Index(i).Set("showAsFlex", wasmQueryShowAsFlex)
		}
		returnValue.Set("css", wasmQueryCSSForArray)
		returnValue.Set("hide", wasmQueryHideForArray)
		returnValue.Set("show", wasmQueryShowForArray)
		returnValue.Set("showAsBlock", wasmQueryShowAsBlockForArray)
		returnValue.Set("showAsInline", wasmQueryShowAsInlineForArray)
		returnValue.Set("showAsInlineBlock", wasmQueryShowAsInlineBlockForArray)
		returnValue.Set("showAsFlex", wasmQueryShowAsFlexForArray)
		return returnValue
	}
	return nil
})
