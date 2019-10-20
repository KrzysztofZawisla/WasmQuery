package mathquery

import (
	"syscall/js"
)

// WasmQueryPhiAsString is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryPhiAsString(this js.Value, args []js.Value) interface{} {
	return js.ValueOf("1.61803398874989484820458683436563811772030917980576286213544862")
}
