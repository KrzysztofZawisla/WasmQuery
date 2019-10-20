package geolocationquery

// WasmQueryReleaseMemory ...
func WasmQueryReleaseMemory() {
	for key, value := range WasmQueryFunctionsStorage {
		value.Release()
		delete(WasmQueryFunctionsStorage, key)
	}
}
