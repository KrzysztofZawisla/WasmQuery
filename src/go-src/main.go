package main

import (
	"syscall/js"

	"github.com/KrzysztofZawisla/WasmQuery/core"
	cryptoquery "github.com/KrzysztofZawisla/WasmQuery/cryptoQuery"
	geolocationquery "github.com/KrzysztofZawisla/WasmQuery/geolocationQuery"
	hashquery "github.com/KrzysztofZawisla/WasmQuery/hashQuery"
	mathquery "github.com/KrzysztofZawisla/WasmQuery/mathQuery"
	securequery "github.com/KrzysztofZawisla/WasmQuery/secureQuery"
	"github.com/KrzysztofZawisla/WasmQuery/shortquery"
	timequery "github.com/KrzysztofZawisla/WasmQuery/timeQuery"
)

func main() {
	c := make(chan struct{}, 0)
	defer close(c)
	shortquery.SetUpFunctionsToVariables()
	defer shortquery.WasmQueryReleaseMemory()
	timequery.SetUpFunctionsToVariables()
	defer timequery.WasmQueryReleaseMemory()
	hashquery.SetUpFunctionsToVariables()
	defer hashquery.WasmQueryReleaseMemory()
	cryptoquery.SetUpFunctionsToVariables()
	defer cryptoquery.WasmQueryReleaseMemory()
	mathquery.SetUpFunctionsToVariables()
	defer mathquery.WasmQueryReleaseMemory()
	geolocationquery.SetUpFunctionsToVariables()
	defer geolocationquery.WasmQueryReleaseMemory()
	securequery.SetUpFunctionsToVariables()
	defer securequery.WasmQueryReleaseMemory()
	WasmQueryRegisterLibrary := js.FuncOf(core.WasmQueryRegisterLibrary)
	js.Global().Set("registerWasmQuery", WasmQueryRegisterLibrary)
	defer WasmQueryRegisterLibrary.Release()
	<-c
}
