package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryMinInt16 is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryMinInt16(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(math.MinInt16)
}
