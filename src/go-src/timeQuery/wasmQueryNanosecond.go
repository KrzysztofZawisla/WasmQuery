package timequery

import (
	"syscall/js"
	"time"
)

// WasmQueryNanosecond is a function from timeQuery module. This function return actual nanosecond.
func WasmQueryNanosecond(this js.Value, args []js.Value) interface{} {
	nanosecond := time.Now().Nanosecond()
	return js.ValueOf(nanosecond)
}
