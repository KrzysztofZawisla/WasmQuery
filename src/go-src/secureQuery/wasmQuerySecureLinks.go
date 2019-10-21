package securequery

import (
	"syscall/js"
)

// WasmQuerySecureLinks is a function from secureQuery module. This function secure all A tags on website.
func WasmQuerySecureLinks(this js.Value, args []js.Value) interface{} {
	elemets := js.Global().Get("document").Call("querySelectorAll", "a")
	patternForNoOpener := js.Global().Get("RegExp").New("noopener", "gmi")
	patternForNoReferrer := js.Global().Get("RegExp").New("noreferrer", "gmi")
	for i := 0; i < elemets.Length(); i++ {
		isEmpty := elemets.Index(i).Get("rel").String() == ""
		if isEmpty {
			elemets.Index(i).Set("rel", js.ValueOf("noopener noreferrer"))
		} else {
			isNoOpener := patternForNoOpener.Call("test", elemets.Index(i).Get("rel")).Bool()
			isNoReferrer := patternForNoReferrer.Call("test", elemets.Index(i).Get("rel")).Bool()
			if !isNoOpener {
				elemets.Index(i).Set("rel", js.ValueOf(elemets.Index(i).Get("rel").String()+" noopener"))
			}
			if !isNoReferrer {
				elemets.Index(i).Set("rel", js.ValueOf(elemets.Index(i).Get("rel").String()+" noreferrer"))
			}
		}
	}
	return js.Undefined()
}
