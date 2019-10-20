package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryPow is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryPow(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0].Float()
		if len(args) > 1 {
			for i := 1; i < len(args); i++ {
				value = math.Pow(value, args[i].Float())
			}
		}
		return js.ValueOf(value)
	}
	return js.Null()
}
