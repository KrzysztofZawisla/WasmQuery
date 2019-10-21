package securequery

import (
	"syscall/js"
)

// WasmQueryRedirectToHTTPS is a function from secureQuery module. This function redirect origin source to https protocol.
func WasmQueryRedirectToHTTPS(this js.Value, args []js.Value) interface{} {
	urlPattern := js.Global().Get("RegExp").New("http://", "gmi")
	url := js.Global().Get("window").Get("location").Get("href")
	url = js.Global().Get("String").New(url)
	isHTTP := urlPattern.Call("test", url).Bool()
	if isHTTP {
		url = url.Call("replace", "http://", "https://")
		js.Global().Get("window").Get("location").Set("href", url)
		return js.ValueOf(true)
	}
	return js.ValueOf(false)
}
