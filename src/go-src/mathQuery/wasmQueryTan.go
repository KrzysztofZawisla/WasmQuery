package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryTan is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryTan(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeNumber {
			return js.ValueOf(math.Tan(value.Float()))
		}
	}
	return js.Null()
}
