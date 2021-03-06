package mathquery

import (
	"math/rand"
	"syscall/js"
	"time"
)

// WasmQueryRandomFloat32 is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryRandomFloat32(this js.Value, args []js.Value) interface{} {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return js.ValueOf(random.Float32())
}
