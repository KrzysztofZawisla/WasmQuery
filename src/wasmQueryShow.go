package main

import "syscall/js"

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
