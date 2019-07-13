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
			return returnValue
		} else if string(domSelector[0]) == "." {
			returnValue := js.Global().Get("document").Call("getElementsByClassName", selectorWithoutSymbol)
			for i := 0; i < returnValue.Length(); i++ {
				returnValue.Index(i).Set("css", wasmQueryCSS)
			}
			returnValue.Set("css", wasmQueryCSSForArray)
			return returnValue
		}
	} else {
		returnValue := js.Global().Get("document").Call("getElementsByTagName", domSelector)
		for i := 0; i < returnValue.Length(); i++ {
			returnValue.Index(i).Set("css", wasmQueryCSS)
		}
		returnValue.Set("css", wasmQueryCSSForArray)
		return returnValue
	}
	return nil
})

var wasmQueryCSS js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return nil
	}
	var selector string = args[0].String()
	if len(args) == 1 {
		return this.Get("style").Get(selector)
	} else if len(args) >= 2 {
		var value string = args[1].String()
		this.Get("style").Set(selector, value)
		return this.Get("style").Get(selector)
	}
	return nil
})

var wasmQueryCSSForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return nil
	}
	var selector string = args[0].String()
	if len(args) >= 2 {
		var value string = args[1].String()
		for i := 0; i < this.Length(); i++ {
			this.Index(i).Get("style").Set(selector, value)
		}
	}
	return nil
})

var disableLibrary js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	wasmQuery.Release()
	wasmQueryCSS.Release()
	wasmQueryCSSForArray.Release()
	return nil
})
