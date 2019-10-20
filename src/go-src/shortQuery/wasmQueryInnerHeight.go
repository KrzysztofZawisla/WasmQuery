package shortquery

import (
	"syscall/js"
)

// WasmQueryInnerHeight is a function from shortQuery module. This function sets or returns the sum of height and padding of html element. Can work only on a single element. Never returns undefined.
func WasmQueryInnerHeight(this js.Value, args []js.Value) interface{} {
	height := js.Global().Get("window").Call("getComputedStyle", this).Get("height").String()
	paddingTop := js.Global().Get("window").Call("getComputedStyle", this).Get("padding-top").String()
	paddingBottom := js.Global().Get("window").Call("getComputedStyle", this).Get("padding-bottom").String()
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeObject && value.Length() > 0 {
			for i := 0; i < value.Length(); i++ {
				if value.Index(i).Type() == js.TypeString {
					this.Get("style").Set("height", "calc("+value.Index(i).String()+" - "+paddingTop+" - "+paddingBottom+")")
				} else if value.Index(i).Type() == js.TypeNumber {
					this.Get("style").Set("height", "calc("+value.Index(i).String()+"px - "+paddingTop+" - "+paddingBottom+")")
				}
			}
			return WasmQueryInnerHeight(this, []js.Value{})

		} else if value.Type() == js.TypeString {
			this.Get("style").Set("height", "calc("+value.String()+" - "+paddingTop+" - "+paddingBottom+")")
			return WasmQueryInnerHeight(this, []js.Value{})

		} else if value.Type() == js.TypeNumber {
			this.Get("style").Set("height", "calc("+value.String()+"px - "+paddingTop+" - "+paddingBottom+")")
			return WasmQueryInnerHeight(this, []js.Value{})
		}
	}
	fHeight := js.Global().Call("parseFloat", js.ValueOf(height[:len(height)-2])).Float()
	fPaddingTop := js.Global().Call("parseFloat", js.ValueOf(paddingTop[:len(paddingTop)-2])).Float()
	fPaddingBottom := js.Global().Call("parseFloat", js.ValueOf(paddingBottom[:len(paddingBottom)-2])).Float()
	return js.Global().Call("String", js.ValueOf(fHeight+fPaddingTop+fPaddingBottom)).String() + "px"
}
