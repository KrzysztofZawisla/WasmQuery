package timequery

import (
	"syscall/js"
	"time"
)

// WasmQueryMinute is a function from timeQuery module. This function return actual minute.
func WasmQueryMinute(this js.Value, args []js.Value) interface{} {
	_, minute, _ := time.Now().Clock()
	return js.ValueOf(minute)
}
