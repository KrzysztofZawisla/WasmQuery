package mathquery

import (
	"syscall/js"
)

// WasmQueryPiAsString is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryPiAsString(this js.Value, args []js.Value) interface{} {
	return js.ValueOf("3.14159265358979323846264338327950288419716939937510582097494459")
}
