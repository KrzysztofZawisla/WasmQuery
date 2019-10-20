package mathquery

import (
	"math"
	"syscall/js"
)

// WasmQueryAcos is a function from mathQuery module. This function returns the arccosine, in radians, of the passed values.
func WasmQueryAcos(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		value := args[0]
		if value.Type() == js.TypeNumber {
			return js.ValueOf(math.Acos(value.Float()))
		} else if value.Type() == js.TypeObject {
			var outputArray []interface{}
			jsOutputArray := js.ValueOf(outputArray)
			for i := 0; i < value.Length(); i++ {
				if value.Index(i).Type() == js.TypeNumber {
					jsOutputArray.SetIndex(i, js.ValueOf(math.Acos(value.Index(i).Float())))
				} else {
					jsOutputArray.SetIndex(i, js.Global().Get("Error").New("Wrongly passed argument. Value have the wrong type: "+value.Index(i).Type().String()+". Expected: "+js.TypeNumber.String()+"."))
				}
			}
			return jsOutputArray
		}
		return js.Global().Get("Error").New("Wrongly passed argument. Value have the wrong type: " + value.Type().String() + ". Expected: " + js.TypeNumber.String() + ".")
	}
	return js.Global().Get("Error").New("Wrongly passed argument. Function require at least 1 argument but got 0.")
}
