package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryPi is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryPi(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(math.Pi)
}
