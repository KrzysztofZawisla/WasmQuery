package core

import (
	"syscall/js"

	"github.com/KrzysztofZawisla/WasmQuery/shortquery"
)

/* WasmQueryDisableLibrary is a function to disable whole library or parts of it.  */
var WasmQueryDisableLibrary js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	arrayWithMethods := make(map[int]interface{})
	arrayWithMethods[0] = shortquery.WasmQuery
	arrayWithMethods[1] = shortquery.WasmQueryCSS
	arrayWithMethods[2] = shortquery.WasmQueryCSSForArray
	arrayWithMethods[3] = shortquery.WasmQueryHide
	arrayWithMethods[4] = shortquery.WasmQueryHideForArray
	arrayWithMethods[5] = shortquery.WasmQueryHideForArray
	arrayWithMethods[6] = shortquery.WasmQueryShow
	arrayWithMethods[7] = shortquery.WasmQueryShowForArray
	arrayWithMethods[8] = shortquery.WasmQueryShowAsBlock
	arrayWithMethods[9] = shortquery.WasmQueryShowAsBlockForArray
	arrayWithMethods[10] = shortquery.WasmQueryShowAsInline
	arrayWithMethods[11] = shortquery.WasmQueryShowAsInlineForArray
	arrayWithMethods[12] = shortquery.WasmQueryShowAsInlineBlock
	arrayWithMethods[13] = shortquery.WasmQueryShowAsInlineBlockForArray
	arrayWithMethods[14] = shortquery.WasmQueryShowAsFlex
	arrayWithMethods[15] = shortquery.WasmQueryShowAsFlexForArray
	arrayWithMethods[16] = shortquery.WasmQueryVal
	arrayWithMethods[17] = shortquery.WasmQueryValForArray
	arrayWithMethods[18] = shortquery.WasmQueryLen
	arrayWithMethods[19] = shortquery.WasmQueryWidth
	arrayWithMethods[20] = shortquery.WasmQueryWidthForArray
	arrayWithMethods[21] = shortquery.WasmQueryHeight
	arrayWithMethods[22] = shortquery.WasmQueryHeightForArray
	arrayWithMethods[23] = shortquery.WasmQueryAttr
	arrayWithMethods[24] = shortquery.WasmQueryAttrForArray
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
