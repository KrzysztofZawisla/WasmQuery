package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryTrunc is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryTrunc(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeNumber {
			return js.ValueOf(math.Trunc(value.Float()))
		}
	}
	return js.Null()
}
