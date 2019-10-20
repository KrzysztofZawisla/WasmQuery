package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryTanh is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryTanh(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeNumber {
			return js.ValueOf(math.Tanh(value.Float()))
		}
	}
	return js.Null()
}
