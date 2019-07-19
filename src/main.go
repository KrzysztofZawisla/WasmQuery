package main

import (
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("registerWasmQuery", wasmQueryRegisterLibrary)
	<-c
}
