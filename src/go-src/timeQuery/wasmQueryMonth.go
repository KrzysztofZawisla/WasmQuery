package timequery

import (
	"syscall/js"
	"time"
)

// WasmQueryMonth is a function from timeQuery module. This function return actual month.
func WasmQueryMonth(this js.Value, args []js.Value) interface{} {
	_, month, _ := time.Now().Date()
	return js.ValueOf(month.String())
}
