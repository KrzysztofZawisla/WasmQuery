package shortquery

import (
	"syscall/js"
)

// WasmQueryOuterHeight is a function from shortQuery module. This function sets or returns the sum of height, padding and border of html element. Can work only on a single element. Never returns undefined.
func WasmQueryOuterHeight(this js.Value, args []js.Value) interface{} {
	height := js.Global().Get("window").Call("getComputedStyle", this).Get("height").String()
	paddingTop := js.Global().Get("window").Call("getComputedStyle", this).Get("padding-top").String()
	paddingBottom := js.Global().Get("window").Call("getComputedStyle", this).Get("padding-bottom").String()
	borderTop := js.Global().Get("window").Call("getComputedStyle", this).Get("border-top-width").String()
	borderBottom := js.Global().Get("window").Call("getComputedStyle", this).Get("border-bottom-width").String()
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeObject && value.Length() > 0 {
			for i := 0; i < value.Length(); i++ {
				if value.Index(i).Type() == js.TypeString {
					this.Get("style").Set("height", "calc("+value.Index(i).String()+" - "+paddingTop+" - "+paddingBottom+" - "+borderTop+" - "+borderBottom+")")
				} else if value.Index(i).Type() == js.TypeNumber {
					this.Get("style").Set("height", "calc("+value.Index(i).String()+"px - "+paddingTop+" - "+paddingBottom+" - "+borderTop+" - "+borderBottom+")")
				}
			}
			return WasmQueryOuterHeight(this, []js.Value{})
		} else if value.Type() == js.TypeString {
			this.Get("style").Set("height", "calc("+value.String()+" - "+paddingTop+" - "+paddingBottom+" - "+borderTop+" - "+borderBottom+")")
			return WasmQueryOuterHeight(this, []js.Value{})
		} else if value.Type() == js.TypeNumber {
			this.Get("style").Set("height", "calc("+value.String()+"px - "+paddingTop+" - "+paddingBottom+" - "+borderTop+" - "+borderBottom+")")
			return WasmQueryOuterHeight(this, []js.Value{})
		}
	}
	fHeight := js.Global().Call("parseFloat", js.ValueOf(height[:len(height)-2])).Float()
	fPaddingTop := js.Global().Call("parseFloat", js.ValueOf(paddingTop[:len(paddingTop)-2])).Float()
	fPaddingBottom := js.Global().Call("parseFloat", js.ValueOf(paddingBottom[:len(paddingBottom)-2])).Float()
	fBorderTop := js.Global().Call("parseFloat", js.ValueOf(borderTop[:len(borderTop)-2])).Float()
	fBorderBottom := js.Global().Call("parseFloat", js.ValueOf(borderBottom[:len(borderBottom)-2])).Float()
	return js.Global().Call("String", js.ValueOf(fHeight+fPaddingTop+fPaddingBottom+fBorderBottom+fBorderTop)).String() + "px"
}
