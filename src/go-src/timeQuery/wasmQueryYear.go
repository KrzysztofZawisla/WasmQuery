package timequery

import (
	"syscall/js"
	"time"
)

// WasmQueryYear is a function from timeQuery module. This function return actual year.
func WasmQueryYear(this js.Value, args []js.Value) interface{} {
	year, _, _ := time.Now().Date()
	return js.ValueOf(year)
}
