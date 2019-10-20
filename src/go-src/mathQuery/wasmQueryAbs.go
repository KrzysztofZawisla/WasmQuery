package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryAbs is a function from mathQuery module. This function returns the absolute of the value.
func WasmQueryAbs(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeNumber {
			return js.ValueOf(math.Abs(value.Float()))
		} else if value.Type() == js.TypeObject {
			var outputArray []interface{}
			jsOutputArray := js.ValueOf(outputArray)
			for i := 0; i < value.Length(); i++ {
				if value.Index(i).Type() == js.TypeNumber {
					jsOutputArray.SetIndex(i, js.ValueOf(math.Abs(value.Index(i).Float())))
				} else {
					jsOutputArray.SetIndex(i, js.Global().Get("Error").New("Wrongly passed argument. Value have the wrong type: "+value.Index(i).Type().String()+". Expected: "+js.TypeNumber.String()+"."))
				}
			}
			return jsOutputArray
		}
		return js.Global().Get("Error").New("Wrongly passed argument. Value have the wrong type: " + value.Type().String() + ". Expected: " + js.TypeNumber.String() + ".")
	}
	return js.Global().Get("Error").New("Wrongly passed argument. Function require 1 argument but got 0.")
}
