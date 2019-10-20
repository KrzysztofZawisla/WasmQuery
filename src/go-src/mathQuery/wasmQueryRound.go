package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryRound is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryRound(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeNumber {
			return js.ValueOf(math.Round(value.Float()))
		}
	}
	return js.Null()
}
