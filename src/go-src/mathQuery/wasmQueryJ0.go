package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryJ0 is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryJ0(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeNumber {
			return js.ValueOf(math.J0(value.Float()))
		}
	}
	return js.Null()
}
