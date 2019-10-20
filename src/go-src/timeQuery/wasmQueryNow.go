package timequery

import (
	"syscall/js"
	"time"
)

// WasmQueryNow is a function from timeQuery module. This function return actual date.
func WasmQueryNow(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(time.Now().String())
}
