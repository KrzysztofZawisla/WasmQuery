package mathquery

import (
	"math/rand"
	"syscall/js"
	"time"
)

// WasmQueryRandomInt64 is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryRandomInt64(this js.Value, args []js.Value) interface{} {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	if len(args) > 0 {
		maxValue := int64(args[0].Int())
		return js.ValueOf(random.Int63n(maxValue))
	}
	return js.ValueOf(random.Int63())
}
