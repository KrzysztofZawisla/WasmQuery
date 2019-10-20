package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQuerySqrt is a function from cryptQuery module. This function return keccak256 hash.
func WasmQuerySqrt(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0].Float()
		value = math.Sqrt(value)
		return js.ValueOf(value)
	}
	return js.Null()
}
