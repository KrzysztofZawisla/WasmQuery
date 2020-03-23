package shortquery

import (
	"syscall/js"
)

// WasmQuery is a constructor function of shortQuery module.
func WasmQuery(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return nil
	}
	domSelector := args[0]
	document := js.Global().Get("document")
	if domSelector.Type() == js.TypeString {
		if domSelector.String() == "document" {
			return document
		}
	}
	returnValue := document.Call("querySelectorAll", domSelector)
	for i := 0; i < returnValue.Length(); i++ {
		returnValue.Index(i).Set("on", WasmQueryFunctionsStorage["WasmQueryOn"])
		returnValue.Index(i).Set("addClass", WasmQueryFunctionsStorage["WasmQueryAddClass"])
		returnValue.Index(i).Set("hasClass", WasmQueryFunctionsStorage["WasmQueryHasClass"])
		returnValue.Index(i).Set("removeClass", WasmQueryFunctionsStorage["WasmQueryRemoveClass"])
		returnValue.Index(i).Set("toggleClass", WasmQueryFunctionsStorage["WasmQueryToggleClass"])
		returnValue.Index(i).Set("val", WasmQueryFunctionsStorage["WasmQueryVal"])
		returnValue.Index(i).Set("attr", WasmQueryFunctionsStorage["WasmQueryAttr"])
		returnValue.Index(i).Set("css", WasmQueryFunctionsStorage["WasmQueryCSS"])
		returnValue.Index(i).Set("removeAttr", WasmQueryFunctionsStorage["WasmQueryRemoveAttr"])
		returnValue.Index(i).Set("html", WasmQueryFunctionsStorage["WasmQueryHTML"])
		returnValue.Index(i).Set("hide", WasmQueryFunctionsStorage["WasmQueryHide"])
		returnValue.Index(i).Set("show", WasmQueryFunctionsStorage["WasmQueryShow"])
		returnValue.Index(i).Set("showAsInitial", WasmQueryFunctionsStorage["WasmQueryShowAsInitial"])
		returnValue.Index(i).Set("width", WasmQueryFunctionsStorage["WasmQueryWidth"])
		returnValue.Index(i).Set("innerWidth", WasmQueryFunctionsStorage["WasmQueryInnerWidth"])
		returnValue.Index(i).Set("outerWidth", WasmQueryFunctionsStorage["WasmQueryOuterWidth"])
		returnValue.Index(i).Set("outerWidthMargin", WasmQueryFunctionsStorage["WasmQueryOuterWidthMargin"])
		returnValue.Index(i).Set("height", WasmQueryFunctionsStorage["WasmQueryHeight"])
		returnValue.Index(i).Set("innerHeight", WasmQueryFunctionsStorage["WasmQueryInnerHeight"])
		returnValue.Index(i).Set("outerHeight", WasmQueryFunctionsStorage["WasmQueryOuterHeight"])
		returnValue.Index(i).Set("outerHeightMargin", WasmQueryFunctionsStorage["WasmQueryOuterHeightMargin"])
		returnValue.Index(i).Set("showAsBlock", WasmQueryFunctionsStorage["WasmQueryShowAsBlock"])
		returnValue.Index(i).Set("showAsInline", WasmQueryFunctionsStorage["WasmQueryShowAsInline"])
		returnValue.Index(i).Set("showAsInlineBlock", WasmQueryFunctionsStorage["WasmQueryShowAsInlineBlock"])
		returnValue.Index(i).Set("showAsFlex", WasmQueryFunctionsStorage["WasmQueryShowAsFlex"])
		returnValue.Index(i).Set("showAsTable", WasmQueryFunctionsStorage["WasmQueryShowAsTable"])
		returnValue.Index(i).Set("showAsGrid", WasmQueryFunctionsStorage["WasmQueryShowAsGrid"])
		returnValue.Index(i).Set("showAsContents", WasmQueryFunctionsStorage["WasmQueryShowAsContents"])
		returnValue.Index(i).Set("showAsInlineFlex", WasmQueryFunctionsStorage["WasmQueryShowAsInlineFlex"])
		returnValue.Index(i).Set("showAsInlineGrid", WasmQueryFunctionsStorage["WasmQueryShowAsInlineGrid"])
		returnValue.Index(i).Set("showAsInlineTable", WasmQueryFunctionsStorage["WasmQueryShowAsInlineTable"])
		returnValue.Index(i).Set("showAsListItem", WasmQueryFunctionsStorage["WasmQueryShowAsListItem"])
		returnValue.Index(i).Set("linkTo", WasmQueryFunctionsStorage["WasmQueryHref"])
	}
	returnValue.Set("addClass", WasmQueryFunctionsStorage["WasmQueryAddClassForArray"])
	returnValue.Set("hasClass", WasmQueryFunctionsStorage["WasmQueryHasClassForArray"])
	returnValue.Set("removeClass", WasmQueryFunctionsStorage["WasmQueryRemoveClassForArray"])
	returnValue.Set("toggleClass", WasmQueryFunctionsStorage["WasmQueryToggleClassForArray"])
	returnValue.Set("val", WasmQueryFunctionsStorage["WasmQueryValForArray"])
	returnValue.Set("attr", WasmQueryFunctionsStorage["WasmQueryAttrForArray"])
	returnValue.Set("css", WasmQueryFunctionsStorage["WasmQueryCSSForArray"])
	returnValue.Set("removeAttr", WasmQueryFunctionsStorage["WasmQueryRemoveAttrForArray"])
	returnValue.Set("width", WasmQueryFunctionsStorage["WasmQueryWidthForArray"])
	returnValue.Set("innerWidth", WasmQueryFunctionsStorage["WasmQueryInnerWidthForArray"])
	returnValue.Set("outerWidth", WasmQueryFunctionsStorage["WasmQueryOuterWidthForArray"])
	returnValue.Set("outerWidthMargin", WasmQueryFunctionsStorage["WasmQueryOuterWidthMarginForArray"])
	returnValue.Set("height", WasmQueryFunctionsStorage["WasmQueryHeightForArray"])
	returnValue.Set("innerHeight", WasmQueryFunctionsStorage["WasmQueryInnerHeightForArray"])
	returnValue.Set("outerHeight", WasmQueryFunctionsStorage["WasmQueryOuterHeightForArray"])
	returnValue.Set("outerHeightMargin", WasmQueryFunctionsStorage["WasmQueryOuterHeightMarginForArray"])
	returnValue.Set("odd", WasmQueryFunctionsStorage["WasmQueryOdd"])
	returnValue.Set("even", WasmQueryFunctionsStorage["WasmQueryEven"])
	returnValue.Set("len", WasmQueryFunctionsStorage["WasmQueryLen"])
	returnValue.Set("html", WasmQueryFunctionsStorage["WasmQueryHTMLForArray"])
	returnValue.Set("hide", WasmQueryFunctionsStorage["WasmQueryHideForArray"])
	returnValue.Set("show", WasmQueryFunctionsStorage["WasmQueryShowForArray"])
	returnValue.Set("showAsInitial", WasmQueryFunctionsStorage["WasmQueryShowAsInitialForArray"])
	returnValue.Set("showAsBlock", WasmQueryFunctionsStorage["WasmQueryShowAsBlockForArray"])
	returnValue.Set("showAsInline", WasmQueryFunctionsStorage["WasmQueryShowAsInlineForArray"])
	returnValue.Set("showAsInlineBlock", WasmQueryFunctionsStorage["WasmQueryShowAsInlineBlockForArray"])
	returnValue.Set("showAsFlex", WasmQueryFunctionsStorage["WasmQueryShowAsFlexForArray"])
	returnValue.Set("showAsTable", WasmQueryFunctionsStorage["WasmQueryShowAsTableForArray"])
	returnValue.Set("showAsGrid", WasmQueryFunctionsStorage["WasmQueryShowAsGridForArray"])
	returnValue.Set("showAsContents", WasmQueryFunctionsStorage["WasmQueryShowAsContentsForArray"])
	returnValue.Set("showAsInlineFlex", WasmQueryFunctionsStorage["WasmQueryShowAsInlineFlexForArray"])
	returnValue.Set("showAsInlineGrid", WasmQueryFunctionsStorage["WasmQueryShowAsInlineGridForArray"])
	returnValue.Set("showAsInlineTable", WasmQueryFunctionsStorage["WasmQueryShowAsInlineTableForArray"])
	returnValue.Set("showAsListItem", WasmQueryFunctionsStorage["WasmQueryShowAsListItemForArray"])
	returnValue.Set("linkTo", WasmQueryFunctionsStorage["WasmQueryHrefForArray"])
	return returnValue
}
