package timequery

import (
	"syscall/js"
	"time"
)

// WasmQueryDay is a function from timeQuery module. This function return actual day.
func WasmQueryDay(this js.Value, args []js.Value) interface{} {
	_, _, day := time.Now().Date()
	return js.ValueOf(day)
}
