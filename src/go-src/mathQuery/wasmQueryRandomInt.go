package mathquery

import (
	"math/rand"
	"syscall/js"
	"time"
)

// WasmQueryRandomInt is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryRandomInt(this js.Value, args []js.Value) interface{} {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	if len(args) > 0 {
		maxValue := args[0].Int()
		return js.ValueOf(random.Intn(maxValue))
	}
	return js.ValueOf(random.Int())
}
