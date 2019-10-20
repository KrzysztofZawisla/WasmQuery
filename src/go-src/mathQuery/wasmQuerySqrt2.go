package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQuerySqrt2 is a function from cryptQuery module. This function return keccak256 hash.
func WasmQuerySqrt2(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(math.Sqrt2)
}
