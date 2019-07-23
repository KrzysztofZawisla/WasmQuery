package timequery

import (
	"syscall/js"
	"time"
)

/* WasmQueryTimeNow is a function from timeQuery module. This function return actual date. */
var WasmQueryTimeNow js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	return time.Now().String()
})
