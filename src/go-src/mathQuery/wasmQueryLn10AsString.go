package mathquery

import (
	"syscall/js"
)

// WasmQueryLn10AsString is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryLn10AsString(this js.Value, args []js.Value) interface{} {
	return js.ValueOf("2.30258509299404568401799145468436420760110148862877297603332790")
}
