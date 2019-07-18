package main

import (
	"syscall/js"
)

var disableLibrary js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	arrayWithMethods := make(map[int]interface{})
	arrayWithMethods[0] = wasmQuery
	arrayWithMethods[1] = wasmQueryCSS
	arrayWithMethods[2] = wasmQueryCSSForArray
	arrayWithMethods[3] = wasmQueryHide
	arrayWithMethods[4] = wasmQueryHideForArray
	arrayWithMethods[5] = wasmQueryHideForArray
	arrayWithMethods[6] = wasmQueryShow
	arrayWithMethods[7] = wasmQueryShowForArray
	arrayWithMethods[8] = wasmQueryShowAsBlock
	arrayWithMethods[9] = wasmQueryShowAsBlockForArray
	arrayWithMethods[10] = wasmQueryShowAsInline
	arrayWithMethods[11] = wasmQueryShowAsInlineForArray
	arrayWithMethods[12] = wasmQueryShowAsInlineBlock
	arrayWithMethods[13] = wasmQueryShowAsInlineBlockForArray
	arrayWithMethods[14] = wasmQueryShowAsFlex
	arrayWithMethods[15] = wasmQueryShowAsFlexForArray
	arrayWithMethods[16] = wasmQueryVal
	arrayWithMethods[17] = wasmQueryValForArray
	arrayWithMethods[18] = wasmQueryLen
	arrayWithMethods[19] = wasmQueryWidth
	arrayWithMethods[20] = wasmQueryWidthForArray
	arrayWithMethods[21] = wasmQueryHeight
	arrayWithMethods[22] = wasmQueryHeightForArray
	arrayWithMethods[23] = wasmQueryAttr
	arrayWithMethods[24] = wasmQueryAttrForArray
	if len(args) == 0 {
		for i := 0; i < len(arrayWithMethods); i++ {
			arrayWithMethods[i].(js.Func).Release()
		}
		return nil
	}
	value := args[0]
	if len(args) >= 1 {
		if value.Type().String() == "object" {
			var outputArray []interface{}
			for i := 0; i < value.Length(); i++ {
				if value.Index(i).Type().String() == "number" {
					for j := 0; j < len(arrayWithMethods); j++ {
						if value.Index(i).Int() == j {
							arrayWithMethods[j].(js.Func).Release()
							outputArray = append(outputArray, j)
							break
						}
					}
				} else {
					return nil
				}
			}
			return outputArray
		} else if value.Type().String() == "number" {
			for i := 0; i < len(arrayWithMethods); i++ {
				if i == value.Int() {
					arrayWithMethods[i].(js.Func).Release()
					return value.Int()
				}
			}
		}
	}
	return nil
})
