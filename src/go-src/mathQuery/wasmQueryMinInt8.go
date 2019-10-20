package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryMinInt8 is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryMinInt8(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(math.MinInt8)
}
