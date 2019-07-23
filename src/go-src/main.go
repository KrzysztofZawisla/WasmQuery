package main

import (
	"syscall/js"

	"github.com/KrzysztofZawisla/WasmQuery/core"
)

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("registerWasmQuery", core.WasmQueryRegisterLibrary)
	<-c
}
