package geolocationquery

import (
	"syscall/js"
)

// WasmQueryLongitude is a function from cryptQuery module. This function return sha512 hash.
func WasmQueryLongitude(this js.Value, args []js.Value) interface{} {
	return js.Global().Get("Promise").New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		promiseArgs := args
		js.Global().Get("navigator").Get("geolocation").Call("getCurrentPosition", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			return promiseArgs[0].Invoke(args[0].Get("coords").Get("longitude"))
		}), js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			return promiseArgs[0].Invoke(args[0])
		}))
		return js.Undefined()
	}))
}
