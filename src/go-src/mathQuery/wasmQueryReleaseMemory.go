package mathquery

// WasmQueryReleaseMemory ...
func WasmQueryReleaseMemory() {
	for key, value := range WasmQueryFunctionsStorage {
		value.Release()
		delete(WasmQueryFunctionsStorage, key)
	}
}
