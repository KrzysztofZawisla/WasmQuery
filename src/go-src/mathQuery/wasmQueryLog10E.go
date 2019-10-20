package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryLog10E is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryLog10E(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(math.Log10E)
}
