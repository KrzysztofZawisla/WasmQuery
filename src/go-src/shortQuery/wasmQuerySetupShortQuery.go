package shortquery

import (
	"syscall/js"
)

// WasmQueryFunctionsStorage stores all functions
var WasmQueryFunctionsStorage = make(map[string]js.Func)

// SetUpFunctionsToVariables register function to case in map
func SetUpFunctionsToVariables() {
	WasmQueryFunctionsStorage["WasmQuery"] = js.FuncOf(WasmQuery)
	WasmQueryFunctionsStorage["WasmQueryOn"] = js.FuncOf(WasmQueryOn)
	WasmQueryFunctionsStorage["WasmQueryAddClass"] = js.FuncOf(WasmQueryAddClass)
	WasmQueryFunctionsStorage["WasmQueryHasClass"] = js.FuncOf(WasmQueryHasClass)
	WasmQueryFunctionsStorage["WasmQueryRemoveClass"] = js.FuncOf(WasmQueryRemoveClass)
	WasmQueryFunctionsStorage["WasmQueryToggleClass"] = js.FuncOf(WasmQueryToggleClass)
	WasmQueryFunctionsStorage["WasmQueryVal"] = js.FuncOf(WasmQueryVal)
	WasmQueryFunctionsStorage["WasmQueryAttr"] = js.FuncOf(WasmQueryAttr)
	WasmQueryFunctionsStorage["WasmQueryCSS"] = js.FuncOf(WasmQueryCSS)
	WasmQueryFunctionsStorage["WasmQueryRemoveAttr"] = js.FuncOf(WasmQueryRemoveAttr)
	WasmQueryFunctionsStorage["WasmQueryHTML"] = js.FuncOf(WasmQueryHTML)
	WasmQueryFunctionsStorage["WasmQueryHide"] = js.FuncOf(WasmQueryHide)
	WasmQueryFunctionsStorage["WasmQueryShow"] = js.FuncOf(WasmQueryShow)
	WasmQueryFunctionsStorage["WasmQueryShowAsInitial"] = js.FuncOf(WasmQueryShowAsInitial)
	WasmQueryFunctionsStorage["WasmQueryWidth"] = js.FuncOf(WasmQueryWidth)
	WasmQueryFunctionsStorage["WasmQueryInnerWidth"] = js.FuncOf(WasmQueryInnerWidth)
	WasmQueryFunctionsStorage["WasmQueryOuterWidth"] = js.FuncOf(WasmQueryOuterWidth)
	WasmQueryFunctionsStorage["WasmQueryOuterWidthMargin"] = js.FuncOf(WasmQueryOuterWidthMargin)
	WasmQueryFunctionsStorage["WasmQueryHeight"] = js.FuncOf(WasmQueryHeight)
	WasmQueryFunctionsStorage["WasmQueryInnerHeight"] = js.FuncOf(WasmQueryInnerHeight)
	WasmQueryFunctionsStorage["WasmQueryOuterHeight"] = js.FuncOf(WasmQueryOuterHeight)
	WasmQueryFunctionsStorage["WasmQueryOuterHeightMargin"] = js.FuncOf(WasmQueryOuterHeightMargin)
	WasmQueryFunctionsStorage["WasmQueryShowAsBlock"] = js.FuncOf(WasmQueryShowAsBlock)
	WasmQueryFunctionsStorage["WasmQueryShowAsInline"] = js.FuncOf(WasmQueryShowAsInline)
	WasmQueryFunctionsStorage["WasmQueryShowAsInlineBlock"] = js.FuncOf(WasmQueryShowAsInlineBlock)
	WasmQueryFunctionsStorage["WasmQueryShowAsFlex"] = js.FuncOf(WasmQueryShowAsFlex)
	WasmQueryFunctionsStorage["WasmQueryShowAsTable"] = js.FuncOf(WasmQueryShowAsTable)
	WasmQueryFunctionsStorage["WasmQueryShowAsGrid"] = js.FuncOf(WasmQueryShowAsGrid)
	WasmQueryFunctionsStorage["WasmQueryShowAsContents"] = js.FuncOf(WasmQueryShowAsContents)
	WasmQueryFunctionsStorage["WasmQueryShowAsInlineFlex"] = js.FuncOf(WasmQueryShowAsInlineFlex)
	WasmQueryFunctionsStorage["WasmQueryShowAsInlineGrid"] = js.FuncOf(WasmQueryShowAsInlineGrid)
	WasmQueryFunctionsStorage["WasmQueryShowAsInlineTable"] = js.FuncOf(WasmQueryShowAsInlineTable)
	WasmQueryFunctionsStorage["WasmQueryShowAsListItem"] = js.FuncOf(WasmQueryShowAsListItem)
	WasmQueryFunctionsStorage["WasmQueryAddClassForArray"] = js.FuncOf(WasmQueryAddClassForArray)
	WasmQueryFunctionsStorage["WasmQueryHasClassForArray"] = js.FuncOf(WasmQueryHasClassForArray)
	WasmQueryFunctionsStorage["WasmQueryRemoveClassForArray"] = js.FuncOf(WasmQueryRemoveClassForArray)
	WasmQueryFunctionsStorage["WasmQueryToggleClassForArray"] = js.FuncOf(WasmQueryToggleClassForArray)
	WasmQueryFunctionsStorage["WasmQueryValForArray"] = js.FuncOf(WasmQueryValForArray)
	WasmQueryFunctionsStorage["WasmQueryAttrForArray"] = js.FuncOf(WasmQueryAttrForArray)
	WasmQueryFunctionsStorage["WasmQueryCSSForArray"] = js.FuncOf(WasmQueryCSSForArray)
	WasmQueryFunctionsStorage["WasmQueryRemoveAttrForArray"] = js.FuncOf(WasmQueryRemoveAttrForArray)
	WasmQueryFunctionsStorage["WasmQueryWidthForArray"] = js.FuncOf(WasmQueryWidthForArray)
	WasmQueryFunctionsStorage["WasmQueryInnerWidthForArray"] = js.FuncOf(WasmQueryInnerWidthForArray)
	WasmQueryFunctionsStorage["WasmQueryOuterWidthForArray"] = js.FuncOf(WasmQueryOuterWidthForArray)
	WasmQueryFunctionsStorage["WasmQueryOuterWidthMarginForArray"] = js.FuncOf(WasmQueryOuterWidthMarginForArray)
	WasmQueryFunctionsStorage["WasmQueryHeightForArray"] = js.FuncOf(WasmQueryHeightForArray)
	WasmQueryFunctionsStorage["WasmQueryInnerHeightForArray"] = js.FuncOf(WasmQueryInnerHeightForArray)
	WasmQueryFunctionsStorage["WasmQueryOuterHeightForArray"] = js.FuncOf(WasmQueryOuterHeightForArray)
	WasmQueryFunctionsStorage["WasmQueryOuterHeightMarginForArray"] = js.FuncOf(WasmQueryOuterHeightMarginForArray)
	WasmQueryFunctionsStorage["WasmQueryOdd"] = js.FuncOf(WasmQueryOdd)
	WasmQueryFunctionsStorage["WasmQueryEven"] = js.FuncOf(WasmQueryEven)
	WasmQueryFunctionsStorage["WasmQueryLen"] = js.FuncOf(WasmQueryLen)
	WasmQueryFunctionsStorage["WasmQueryHTMLForArray"] = js.FuncOf(WasmQueryHTMLForArray)
	WasmQueryFunctionsStorage["WasmQueryHideForArray"] = js.FuncOf(WasmQueryHideForArray)
	WasmQueryFunctionsStorage["WasmQueryShowForArray"] = js.FuncOf(WasmQueryShowForArray)
	WasmQueryFunctionsStorage["WasmQueryShowAsInitialForArray"] = js.FuncOf(WasmQueryShowAsInitialForArray)
	WasmQueryFunctionsStorage["WasmQueryShowAsBlockForArray"] = js.FuncOf(WasmQueryShowAsBlockForArray)
	WasmQueryFunctionsStorage["WasmQueryShowAsInlineForArray"] = js.FuncOf(WasmQueryShowAsInlineForArray)
	WasmQueryFunctionsStorage["WasmQueryShowAsInlineBlockForArray"] = js.FuncOf(WasmQueryShowAsInlineBlockForArray)
	WasmQueryFunctionsStorage["WasmQueryShowAsFlexForArray"] = js.FuncOf(WasmQueryShowAsFlexForArray)
	WasmQueryFunctionsStorage["WasmQueryShowAsTableForArray"] = js.FuncOf(WasmQueryShowAsTableForArray)
	WasmQueryFunctionsStorage["WasmQueryShowAsGridForArray"] = js.FuncOf(WasmQueryShowAsGridForArray)
	WasmQueryFunctionsStorage["WasmQueryShowAsContentsForArray"] = js.FuncOf(WasmQueryShowAsContentsForArray)
	WasmQueryFunctionsStorage["WasmQueryShowAsInlineFlexForArray"] = js.FuncOf(WasmQueryShowAsInlineFlexForArray)
	WasmQueryFunctionsStorage["WasmQueryShowAsInlineGridForArray"] = js.FuncOf(WasmQueryShowAsInlineGridForArray)
	WasmQueryFunctionsStorage["WasmQueryShowAsInlineTableForArray"] = js.FuncOf(WasmQueryShowAsInlineTableForArray)
	WasmQueryFunctionsStorage["WasmQueryShowAsListItemForArray"] = js.FuncOf(WasmQueryShowAsListItemForArray)
	WasmQueryFunctionsStorage["WasmQueryHref"] = js.FuncOf(WasmQueryHref)
	WasmQueryFunctionsStorage["WasmQueryHrefForArray"] = js.FuncOf(WasmQueryHrefForArray)
}
