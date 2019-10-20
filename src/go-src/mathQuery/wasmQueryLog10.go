package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryLog10 is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryLog10(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeNumber {
			return js.ValueOf(math.Log10(value.Float()))
		}
	}
	return js.Null()
}
