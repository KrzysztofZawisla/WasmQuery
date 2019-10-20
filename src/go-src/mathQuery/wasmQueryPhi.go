package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryPhi is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryPhi(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(math.Phi)
}
