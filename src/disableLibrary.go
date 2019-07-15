package main

import "syscall/js"

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
	wasmQueryVal.Release()
	wasmQueryValForArray.Release()
	wasmQueryLen.Release()
	return nil
})
