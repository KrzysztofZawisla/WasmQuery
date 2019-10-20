package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryMaxInt8 is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryMaxInt8(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(math.MaxInt8)
}
