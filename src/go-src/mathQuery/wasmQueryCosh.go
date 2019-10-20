package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryCosh is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryCosh(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeNumber {
			return js.ValueOf(math.Cosh(value.Float()))
		}
	}
	return js.Null()
}
