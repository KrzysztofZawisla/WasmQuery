package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQuerySqrtPi is a function from cryptQuery module. This function return keccak256 hash.
func WasmQuerySqrtPi(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(math.SqrtPi)
}
