package timequery

import (
	"syscall/js"
	"time"
)

// WasmQuerySecond is a function from timeQuery module. This function return actual second.
func WasmQuerySecond(this js.Value, args []js.Value) interface{} {
	_, _, second := time.Now().Clock()
	return js.ValueOf(second)
}
