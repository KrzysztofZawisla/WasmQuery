package shortquery

import "syscall/js"

/* WasmQueryShowForArray is a function from shortQuery module. This function set/return display block/inline by smart checking for html elements. This function can only work on arrays of html elements. */
var WasmQueryShowForArray js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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
		valueOfIteration := js.Global().Get("window").Call("getComputedStyle", this.Index(i)).Get("display")
		outputArray = append(outputArray, valueOfIteration)
	}
	return outputArray
})
