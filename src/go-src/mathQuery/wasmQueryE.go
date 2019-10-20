package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryE is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryE(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(math.E)
}
