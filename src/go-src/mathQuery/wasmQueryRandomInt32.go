package mathquery

import (
	"math/rand"
	"syscall/js"
	"time"
)

// WasmQueryRandomInt32 is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryRandomInt32(this js.Value, args []js.Value) interface{} {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	if len(args) > 0 {
		maxValue := int32(args[0].Int())
		return js.ValueOf(random.Int31n(maxValue))
	}
	return js.ValueOf(random.Int31())
}
