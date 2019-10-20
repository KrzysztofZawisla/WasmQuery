package hashquery

import (
	"encoding/hex"
	"syscall/js"

	"hash/fnv"
)

// WasmQueryFNV1a64 is a function from cryptQuery module. This function return keccak256 hash.
func WasmQueryFNV1a64(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		crypto := fnv.New64a()
		crypto.Write([]byte(args[0].String()))
		value := hex.EncodeToString(crypto.Sum(nil))
		return js.ValueOf(value)
	}
	return js.Null()
}
