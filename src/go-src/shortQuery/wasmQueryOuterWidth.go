package shortquery

import (
	"syscall/js"
)

// WasmQueryOuterWidth is a function from shortQuery module. This function sets or returns the sum of width, padding and border of an html element. Can work only on a single element. Never returns undefined.
func WasmQueryOuterWidth(this js.Value, args []js.Value) interface{} {
	width := js.Global().Get("window").Call("getComputedStyle", this).Get("width").String()
	paddingLeft := js.Global().Get("window").Call("getComputedStyle", this).Get("padding-left").String()
	paddingRight := js.Global().Get("window").Call("getComputedStyle", this).Get("padding-right").String()
	borderLeft := js.Global().Get("window").Call("getComputedStyle", this).Get("border-left-width").String()
	borderRight := js.Global().Get("window").Call("getComputedStyle", this).Get("border-right-width").String()
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeObject && value.Length() > 0 {
			for i := 0; i < value.Length(); i++ {
				if value.Index(i).Type() == js.TypeString {
					this.Get("style").Set("width", "calc("+value.Index(i).String()+" - "+paddingLeft+" - "+paddingRight+" - "+borderLeft+" - "+borderRight+")")
				} else if value.Index(i).Type() == js.TypeNumber {
					this.Get("style").Set("width", "calc("+value.Index(i).String()+"px - "+paddingLeft+" - "+paddingRight+" - "+borderLeft+" - "+borderRight+")")
				}
			}
			return WasmQueryOuterWidth(this, []js.Value{})
		} else if value.Type() == js.TypeString {
			this.Get("style").Set("width", "calc("+value.String()+" - "+paddingLeft+" - "+paddingRight+" - "+borderLeft+" - "+borderRight+")")
			return WasmQueryOuterWidth(this, []js.Value{})

		} else if value.Type() == js.TypeNumber {
			this.Get("style").Set("width", "calc("+value.String()+"px - "+paddingLeft+" - "+paddingRight+" - "+borderLeft+" - "+borderRight+")")
			return WasmQueryOuterWidth(this, []js.Value{})
		}
	}
	fHeight := js.Global().Call("parseFloat", js.ValueOf(width[:len(width)-2])).Float()
	fPaddingLeft := js.Global().Call("parseFloat", js.ValueOf(paddingLeft[:len(paddingLeft)-2])).Float()
	fPaddingRight := js.Global().Call("parseFloat", js.ValueOf(paddingRight[:len(paddingRight)-2])).Float()
	fBorderLeft := js.Global().Call("parseFloat", js.ValueOf(borderLeft[:len(borderLeft)-2])).Float()
	fBorderRight := js.Global().Call("parseFloat", js.ValueOf(borderRight[:len(borderRight)-2])).Float()
	return js.Global().Call("String", js.ValueOf(fHeight+fPaddingLeft+fPaddingRight+fBorderRight+fBorderLeft)).String() + "px"
}
