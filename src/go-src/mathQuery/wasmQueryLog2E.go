package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryLog2E is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryLog2E(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(math.Log2E)
}
