package mathquery

import (
	"syscall/js"
)

// WasmQueryLn2AsString is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryLn2AsString(this js.Value, args []js.Value) interface{} {
	return js.ValueOf("0.693147180559945309417232121458176568075500134360255254120680009")
}
