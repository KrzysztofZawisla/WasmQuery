package mathquery

import (
	"syscall/js"
)

// WasmQuerySqrt2AsString is a function from cryptQuery module. This function return keccak256 hash.
func WasmQuerySqrt2AsString(this js.Value, args []js.Value) interface{} {
	return js.ValueOf("1.41421356237309504880168872420969807856967187537694807317667974")
}
