package mathquery

import (
	"syscall/js"
)

// WasmQuerySqrtPiAsString is a function from cryptQuery module. This function return keccak256 hash.
func WasmQuerySqrtPiAsString(this js.Value, args []js.Value) interface{} {
	return js.ValueOf("1.77245385090551602729816748334114518279754945612238712821380779")
}
