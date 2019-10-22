package hashquery

import (
	"encoding/hex"
	"syscall/js"

	"hash/fnv"
)

// WasmQueryFNV1_32 is a function from cryptQuery module. This function return fnv1 32 hash.
func WasmQueryFNV1_32(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeString {
			crypto := fnv.New32()
			crypto.Write([]byte(value.String()))
			value := hex.EncodeToString(crypto.Sum(nil))
			return js.ValueOf(value)
		} else if value.Type() == js.TypeObject {
			var outputArray []interface{}
			jsOutputArray := js.ValueOf(outputArray)
			for i := 0; i < value.Length(); i++ {
				if value.Index(i).Type() == js.TypeString {
					crypto := fnv.New32()
					crypto.Write([]byte(value.Index(i).String()))
					value := hex.EncodeToString(crypto.Sum(nil))
					jsOutputArray.SetIndex(i, js.ValueOf(value))
				} else {
					jsOutputArray.SetIndex(i, js.Global().Get("Error").New("Wrongly passed argument. Value have the wrong type: "+value.Index(i).Type().String()+". Expected: "+js.TypeString.String()+"."))
				}
			}
		}
		return js.Global().Get("Error").New("Wrongly passed argument. Value have the wrong type: " + value.Type().String() + ". Expected: " + js.TypeString.String() + ".")
	}
	return js.Global().Get("Error").New("Wrongly passed argument. Function require 1 argument but got 0.")
}
