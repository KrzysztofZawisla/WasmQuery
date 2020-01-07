package cryptoquery

import (
	"crypto/md5"
	"encoding/hex"
	"syscall/js"
)

// WasmQueryMD5 is a function from cryptQuery module. This function return md5 hash.
func WasmQueryMD5(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeString {
			crypto := md5.New()
			crypto.Write([]byte(args[0].String()))
			value := hex.EncodeToString(crypto.Sum(nil))
			return js.ValueOf(value)
		}
		return js.Global().Get("WasmQueryError").New("Wrongly passed argument. Value have the wrong type: " + value.Type().String() + ". Expected: " + js.TypeString.String() + ".")
	}
	return js.Global().Get("WasmQueryError").New("Wrongly passed argument. Function require 1 argument but got 0.")
}
