package timequery

import (
	"syscall/js"
	"time"
)

// WasmQueryHour is a function from timeQuery module. This function return actual hour.
func WasmQueryHour(this js.Value, args []js.Value) interface{} {
	hour, _, _ := time.Now().Clock()
	return js.ValueOf(hour)
}
