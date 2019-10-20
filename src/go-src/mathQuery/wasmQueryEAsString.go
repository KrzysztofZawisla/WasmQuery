package mathquery

import (
	"syscall/js"
)

// WasmQueryEAsString is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryEAsString(this js.Value, args []js.Value) interface{} {
	return js.ValueOf("2.71828182845904523536028747135266249775724709369995957496696763")
}
