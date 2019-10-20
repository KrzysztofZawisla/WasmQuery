package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryHypot is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryHypot(this js.Value, args []js.Value) interface{} {
	if len(args) > 1 {
		value1 := args[0]
		value2 := args[1]
		if value1.Type() == js.TypeNumber && value2.Type() == js.TypeNumber {
			return js.ValueOf(math.Hypot(value1.Float(), value2.Float()))
		}
	}
	return js.Null()
}
