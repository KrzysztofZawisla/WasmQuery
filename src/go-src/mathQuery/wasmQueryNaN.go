package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryNaN is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryNaN(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(math.NaN())
}
