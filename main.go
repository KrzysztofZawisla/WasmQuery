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
	var outputArray []interface{}
	if len(args) == 1 {
		for i := 0; i < this.Length(); i++ {
			valueOfIteration := this.Index(i).Get("style").Get(selector)
			outputArray = append(outputArray, valueOfIteration)
		}
		return outputArray
	} else if len(args) >= 2 {
		var value string = args[1].String()
		for i := 0; i < this.Length(); i++ {
			this.Index(i).Get("style").Set(selector, value)
			valueOfIteration := this.Index(i).Get("style").Get(selector)
			outputArray = append(outputArray, valueOfIteration)
		}
		return outputArray
	}
	return nil
})

var wasmQueryHide js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	this.Get("style").Set("display", "none")
	return this.Get("style").Get("display")
})

var wasmQueryHideForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	var outputArray []interface{}
	for i := 0; i < this.Length(); i++ {
		this.Index(i).Get("style").Set("display", "none")
		valueOfIteration := this.Index(i).Get("style").Get("display")
		outputArray = append(outputArray, valueOfIteration)
	}
	return outputArray
})

var wasmQueryShow js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	listOfInlineTags := []string{"B", "BIG", "I", "SMALL", "TT", "ABBR", "ACRONYM", "CITE", "CODE", "DFN", "EM", "KBD", "STRONG", "SAMP", "VAR", "A", "BDO", "BR", "IMG", "MAP", "OBJECT", "Q", "SCRIPT", "SPAN", "SUB", "SUP", "BUTTON", "INPUT", "LABEL", "SELECT", "TEXTAREA"}
	for i := 0; i < len(listOfInlineTags); i++ {
		if this.Get("tagName").String() == listOfInlineTags[i] {
			this.Get("style").Set("display", "inline")
			return this.Get("style").Get("display")
		}
	}
	this.Get("style").Set("display", "block")
	return this.Get("style").Get("display")
})

var wasmQueryShowForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	var outputArray []interface{}
	listOfInlineTags := []string{"B", "BIG", "I", "SMALL", "TT", "ABBR", "ACRONYM", "CITE", "CODE", "DFN", "EM", "KBD", "STRONG", "SAMP", "VAR", "A", "BDO", "BR", "IMG", "MAP", "OBJECT", "Q", "SCRIPT", "SPAN", "SUB", "SUP", "BUTTON", "INPUT", "LABEL", "SELECT", "TEXTAREA"}
	for i := 0; i < this.Length(); i++ {
		for j := 0; j < len(listOfInlineTags); j++ {
			if this.Index(i).Get("tagName").String() == listOfInlineTags[j] {
				this.Index(i).Get("style").Set("display", "inline")
				break
			} else {
				this.Index(i).Get("style").Set("display", "block")
			}
		}
		valueOfIteration := this.Index(i).Get("style").Get("display")
		outputArray = append(outputArray, valueOfIteration)
	}
	return outputArray
})

var wasmQueryShowAsBlock js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	this.Get("style").Set("display", "block")
	return this.Get("style").Get("display")
})

var wasmQueryShowAsBlockForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	var outputArray []interface{}
	for i := 0; i < this.Length(); i++ {
		this.Index(i).Get("style").Set("display", "block")
		valueOfIteration := this.Index(i).Get("style").Get("display")
		outputArray = append(outputArray, valueOfIteration)
	}
	return outputArray
})

var wasmQueryShowAsInline js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	this.Get("style").Set("display", "inline")
	return this.Get("style").Get("display")
})

var wasmQueryShowAsInlineForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	var outputArray []interface{}
	for i := 0; i < this.Length(); i++ {
		this.Index(i).Get("style").Set("display", "inline")
		valueOfIteration := this.Index(i).Get("style").Get("display")
		outputArray = append(outputArray, valueOfIteration)
	}
	return outputArray
})

var wasmQueryShowAsInlineBlock js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	this.Get("style").Set("display", "inline-block")
	return this.Get("style").Get("display")
})

var wasmQueryShowAsInlineBlockForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	var outputArray []interface{}
	for i := 0; i < this.Length(); i++ {
		this.Index(i).Get("style").Set("display", "inline-block")
		valueOfIteration := this.Index(i).Get("style").Get("display")
		outputArray = append(outputArray, valueOfIteration)
	}
	return outputArray
})

var wasmQueryShowAsFlex js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	this.Get("style").Set("display", "flex")
	return this.Get("style").Get("display")
})

var wasmQueryShowAsFlexForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	var outputArray []interface{}
	for i := 0; i < this.Length(); i++ {
		this.Index(i).Get("style").Set("display", "flex")
		valueOfIteration := this.Index(i).Get("style").Get("display")
		outputArray = append(outputArray, valueOfIteration)
	}
	return outputArray
})

var disableLibrary js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	wasmQuery.Release()
	wasmQueryCSS.Release()
	wasmQueryCSSForArray.Release()
	wasmQueryHide.Release()
	wasmQueryHideForArray.Release()
	wasmQueryShow.Release()
	wasmQueryShowForArray.Release()
	wasmQueryShowAsBlock.Release()
	wasmQueryShowAsBlockForArray.Release()
	wasmQueryShowAsInline.Release()
	wasmQueryShowAsInlineForArray.Release()
	wasmQueryShowAsInlineBlock.Release()
	wasmQueryShowAsInlineBlockForArray.Release()
	wasmQueryShowAsFlex.Release()
	wasmQueryShowAsFlexForArray.Release()
	return nil
})
