package hashquery

import (
	"encoding/hex"
	"syscall/js"

	"hash/fnv"
)

// WasmQueryFNV1_128 is a function from cryptQuery module. This function return fnv1 128 hash.
func WasmQueryFNV1_128(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeString {
			crypto := fnv.New128()
			crypto.Write([]byte(value.String()))
			value := hex.EncodeToString(crypto.Sum(nil))
			return js.ValueOf(value)
		}
		return js.Global().Get("Error").New("Wrongly passed argument. Value have the wrong type: " + value.Type().String() + ". Expected: " + js.TypeString.String() + ".")
	}
	return js.Global().Get("Error").New("Wrongly passed argument. Function require 1 argument but got 0.")
}
