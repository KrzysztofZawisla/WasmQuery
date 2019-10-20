package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryFloor is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryFloor(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeNumber {
			return js.ValueOf(math.Floor(value.Float()))
		}
	}
	return js.Null()
}
