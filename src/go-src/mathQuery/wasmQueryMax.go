package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryMax is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryMax(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		biggestValue := args[0].Float()
		if len(args) > 1 {
			for i := 1; i < len(args); i++ {
				biggestValue = math.Max(biggestValue, args[i].Float())
			}
		}
		return js.ValueOf(biggestValue)
	}
	return js.Null()
}
