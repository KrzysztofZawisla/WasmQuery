package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryMin is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryMin(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		smallerValue := args[0].Float()
		if len(args) > 1 {
			for i := 1; i < len(args); i++ {
				smallerValue = math.Min(smallerValue, args[i].Float())
			}
		}
		return js.ValueOf(smallerValue)
	}
	return js.Null()
}
