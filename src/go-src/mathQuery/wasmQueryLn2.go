package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryLn2 is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryLn2(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(math.Ln2)
}
