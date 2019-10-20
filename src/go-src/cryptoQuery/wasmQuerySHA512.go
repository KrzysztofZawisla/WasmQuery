package cryptoquery

import (
	"crypto/sha512"
	"encoding/hex"
	"syscall/js"
)

// WasmQuerySHA512 is a function from cryptQuery module. This function return sha512 hash.
func WasmQuerySHA512(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeString {
			crypto := sha512.New()
			crypto.Write([]byte(args[0].String()))
			value := hex.EncodeToString(crypto.Sum(nil))
			return js.ValueOf(value)
		}
		return js.Global().Get("Error").New("Wrongly passed argument. Value have the wrong type: " + value.Type().String() + ". Expected: " + js.TypeString.String() + ".")
	}
	return js.Global().Get("Error").New("Wrongly passed argument. Function require 1 argument but got 0.")
}
