package core

import (
	"syscall/js"

	cryptoquery "github.com/KrzysztofZawisla/WasmQuery/cryptoQuery"
	geolocationquery "github.com/KrzysztofZawisla/WasmQuery/geolocationQuery"
	hashquery "github.com/KrzysztofZawisla/WasmQuery/hashQuery"
	mathquery "github.com/KrzysztofZawisla/WasmQuery/mathQuery"
	securequery "github.com/KrzysztofZawisla/WasmQuery/secureQuery"
	shortquery "github.com/KrzysztofZawisla/WasmQuery/shortquery"
	timequery "github.com/KrzysztofZawisla/WasmQuery/timeQuery"
)

// WasmQueryRegisterLibrary is a function to register library.
func WasmQueryRegisterLibrary(this js.Value, args []js.Value) interface{} {
	symbol := "$"
	emptyObject := make(map[string]interface{})
	if len(args) > 0 {
		if args[0].Type() == js.TypeString {
			symbol = args[0].String()
		} else if args[0].Type() == js.TypeObject {
			if args[0].Index(0).Type() == js.TypeString {
				symbol = args[0].Index(0).String()
			}
		}
	}
	js.Global().Set(symbol, shortquery.WasmQueryFunctionsStorage["WasmQuery"])
	symbolJsObject := js.Global().Get(symbol)
	symbolJsObject.Set("time", js.ValueOf(emptyObject))
	symbolJsObject.Get("time").Set("now", timequery.WasmQueryFunctionsStorage["WasmQueryNow"])
	symbolJsObject.Get("time").Set("month", timequery.WasmQueryFunctionsStorage["WasmQueryMonth"])
	symbolJsObject.Get("time").Set("year", timequery.WasmQueryFunctionsStorage["WasmQueryYear"])
	symbolJsObject.Get("time").Set("day", timequery.WasmQueryFunctionsStorage["WasmQueryDay"])
	symbolJsObject.Get("time").Set("hour", timequery.WasmQueryFunctionsStorage["WasmQueryHour"])
	symbolJsObject.Get("time").Set("minute", timequery.WasmQueryFunctionsStorage["WasmQueryMinute"])
	symbolJsObject.Get("time").Set("second", timequery.WasmQueryFunctionsStorage["WasmQuerySecond"])
	symbolJsObject.Get("time").Set("nanosecond", timequery.WasmQueryFunctionsStorage["WasmQueryNanosecond"])
	symbolJsObject.Set("crypto", js.ValueOf(emptyObject))
	symbolJsObject.Get("crypto").Set("md4", cryptoquery.WasmQueryFunctionsStorage["WasmQueryMD4"])
	symbolJsObject.Get("crypto").Set("md5", cryptoquery.WasmQueryFunctionsStorage["WasmQueryMD5"])
	symbolJsObject.Get("crypto").Set("sha1", cryptoquery.WasmQueryFunctionsStorage["WasmQuerySHA1"])
	symbolJsObject.Get("crypto").Set("sha512_224", cryptoquery.WasmQueryFunctionsStorage["WasmQuerySHA512_224"])
	symbolJsObject.Get("crypto").Set("sha512_256", cryptoquery.WasmQueryFunctionsStorage["WasmQuerySHA512_256"])
	symbolJsObject.Get("crypto").Set("sha256", cryptoquery.WasmQueryFunctionsStorage["WasmQuerySHA256"])
	symbolJsObject.Get("crypto").Set("sha224", cryptoquery.WasmQueryFunctionsStorage["WasmQuerySHA224"])
	symbolJsObject.Get("crypto").Set("sha384", cryptoquery.WasmQueryFunctionsStorage["WasmQuerySHA384"])
	symbolJsObject.Get("crypto").Set("sha512", cryptoquery.WasmQueryFunctionsStorage["WasmQuerySHA512"])
	symbolJsObject.Get("crypto").Set("sha3_256", cryptoquery.WasmQueryFunctionsStorage["WasmQuerySHA3_256"])
	symbolJsObject.Get("crypto").Set("sha3_224", cryptoquery.WasmQueryFunctionsStorage["WasmQuerySHA3_224"])
	symbolJsObject.Get("crypto").Set("sha3_384", cryptoquery.WasmQueryFunctionsStorage["WasmQuerySHA3_384"])
	symbolJsObject.Get("crypto").Set("sha3_512", cryptoquery.WasmQueryFunctionsStorage["WasmQuerySHA3_512"])
	symbolJsObject.Get("crypto").Set("keccak512", cryptoquery.WasmQueryFunctionsStorage["WasmQueryKeccak512"])
	symbolJsObject.Get("crypto").Set("keccak256", cryptoquery.WasmQueryFunctionsStorage["WasmQueryKeccak256"])
	symbolJsObject.Set("hash", js.ValueOf(emptyObject))
	symbolJsObject.Get("hash").Set("adler32", hashquery.WasmQueryFunctionsStorage["WasmQueryAdler32"])
	symbolJsObject.Get("hash").Set("fnv1_128", hashquery.WasmQueryFunctionsStorage["WasmQueryFNV1_128"])
	symbolJsObject.Get("hash").Set("fnv1a_128", hashquery.WasmQueryFunctionsStorage["WasmQueryFNV1a128"])
	symbolJsObject.Get("hash").Set("fnv1_32", hashquery.WasmQueryFunctionsStorage["WasmQueryFNV1_32"])
	symbolJsObject.Get("hash").Set("fnv1a_32", hashquery.WasmQueryFunctionsStorage["WasmQueryFNV1a32"])
	symbolJsObject.Get("hash").Set("fnv1_64", hashquery.WasmQueryFunctionsStorage["WasmQueryFNV1_64"])
	symbolJsObject.Get("hash").Set("fnv1a_64", hashquery.WasmQueryFunctionsStorage["WasmQueryFNV1a64"])
	symbolJsObject.Set("math", js.ValueOf(emptyObject))
	symbolJsObject.Get("math").Set("randomInt", mathquery.WasmQueryFunctionsStorage["WasmQueryRandomInt"])
	symbolJsObject.Get("math").Set("randomInt32", mathquery.WasmQueryFunctionsStorage["WasmQueryRandomInt32"])
	symbolJsObject.Get("math").Set("randomInt64", mathquery.WasmQueryFunctionsStorage["WasmQueryRandomInt64"])
	symbolJsObject.Get("math").Set("randomFloat32", mathquery.WasmQueryFunctionsStorage["WasmQueryRandomFloat32"])
	symbolJsObject.Get("math").Set("randomFloat64", mathquery.WasmQueryFunctionsStorage["WasmQueryRandomFloat64"])
	symbolJsObject.Get("math").Set("ceil", mathquery.WasmQueryFunctionsStorage["WasmQueryCeil"])
	symbolJsObject.Get("math").Set("floor", mathquery.WasmQueryFunctionsStorage["WasmQueryFloor"])
	symbolJsObject.Get("math").Set("round", mathquery.WasmQueryFunctionsStorage["WasmQueryRound"])
	symbolJsObject.Get("math").Set("gamma", mathquery.WasmQueryFunctionsStorage["WasmQueryGamma"])
	symbolJsObject.Get("math").Set("NaN", mathquery.WasmQueryFunctionsStorage["WasmQueryNaN"])
	symbolJsObject.Get("math").Set("max", mathquery.WasmQueryFunctionsStorage["WasmQueryMax"])
	symbolJsObject.Get("math").Set("min", mathquery.WasmQueryFunctionsStorage["WasmQueryMin"])
	symbolJsObject.Get("math").Set("pow", mathquery.WasmQueryFunctionsStorage["WasmQueryPow"])
	symbolJsObject.Get("math").Set("sqrt", mathquery.WasmQueryFunctionsStorage["WasmQuerySqrt"])
	symbolJsObject.Get("math").Set("pi", mathquery.WasmQueryFunctionsStorage["WasmQueryPi"])
	symbolJsObject.Get("math").Set("piAsString", mathquery.WasmQueryFunctionsStorage["WasmQueryPiAsString"])
	symbolJsObject.Get("math").Set("e", mathquery.WasmQueryFunctionsStorage["WasmQueryE"])
	symbolJsObject.Get("math").Set("eAsString", mathquery.WasmQueryFunctionsStorage["WasmQueryEAsString"])
	symbolJsObject.Get("math").Set("phi", mathquery.WasmQueryFunctionsStorage["WasmQueryPhi"])
	symbolJsObject.Get("math").Set("phiAsString", mathquery.WasmQueryFunctionsStorage["WasmQueryPhiAsString"])
	symbolJsObject.Get("math").Set("ln2", mathquery.WasmQueryFunctionsStorage["WasmQueryLn2"])
	symbolJsObject.Get("math").Set("ln2AsString", mathquery.WasmQueryFunctionsStorage["WasmQueryLn2AsString"])
	symbolJsObject.Get("math").Set("ln10", mathquery.WasmQueryFunctionsStorage["WasmQueryLn10"])
	symbolJsObject.Get("math").Set("ln10AsString", mathquery.WasmQueryFunctionsStorage["WasmQueryLn10AsString"])
	symbolJsObject.Get("math").Set("log2E", mathquery.WasmQueryFunctionsStorage["WasmQueryLog2E"])
	symbolJsObject.Get("math").Set("log10E", mathquery.WasmQueryFunctionsStorage["WasmQueryLog10E"])
	symbolJsObject.Get("math").Set("sqrt2", mathquery.WasmQueryFunctionsStorage["WasmQuerySqrt2"])
	symbolJsObject.Get("math").Set("sqrt2AsString", mathquery.WasmQueryFunctionsStorage["WasmQuerySqrt2AsString"])
	symbolJsObject.Get("math").Set("maxInt8", mathquery.WasmQueryFunctionsStorage["WasmQueryMaxInt8"])
	symbolJsObject.Get("math").Set("minInt8", mathquery.WasmQueryFunctionsStorage["WasmQueryMinInt8"])
	symbolJsObject.Get("math").Set("maxInt16", mathquery.WasmQueryFunctionsStorage["WasmQueryMaxInt16"])
	symbolJsObject.Get("math").Set("minInt16", mathquery.WasmQueryFunctionsStorage["WasmQueryMinInt16"])
	symbolJsObject.Get("math").Set("maxInt32", mathquery.WasmQueryFunctionsStorage["WasmQueryMaxInt32"])
	symbolJsObject.Get("math").Set("minInt32", mathquery.WasmQueryFunctionsStorage["WasmQueryMinInt32"])
	symbolJsObject.Get("math").Set("sqrtPi", mathquery.WasmQueryFunctionsStorage["WasmQuerySqrtPi"])
	symbolJsObject.Get("math").Set("sqrtPiAsString", mathquery.WasmQueryFunctionsStorage["WasmQuerySqrtPiAsString"])
	symbolJsObject.Get("math").Set("isNaN", mathquery.WasmQueryFunctionsStorage["WasmQueryIsNaN"])
	symbolJsObject.Get("math").Set("j0", mathquery.WasmQueryFunctionsStorage["WasmQueryJ0"])
	symbolJsObject.Get("math").Set("j1", mathquery.WasmQueryFunctionsStorage["WasmQueryJ1"])
	symbolJsObject.Get("math").Set("log", mathquery.WasmQueryFunctionsStorage["WasmQueryLog"])
	symbolJsObject.Get("math").Set("cosh", mathquery.WasmQueryFunctionsStorage["WasmQueryCosh"])
	symbolJsObject.Get("math").Set("inf", mathquery.WasmQueryFunctionsStorage["WasmQueryInf"])
	symbolJsObject.Get("math").Set("log10", mathquery.WasmQueryFunctionsStorage["WasmQueryLog10"])
	symbolJsObject.Get("math").Set("erf", mathquery.WasmQueryFunctionsStorage["WasmQueryErf"])
	symbolJsObject.Get("math").Set("erfinv", mathquery.WasmQueryFunctionsStorage["WasmQueryErfinv"])
	symbolJsObject.Get("math").Set("erfc", mathquery.WasmQueryFunctionsStorage["WasmQueryErfc"])
	symbolJsObject.Get("math").Set("erfcinv", mathquery.WasmQueryFunctionsStorage["WasmQueryErfcinv"])
	symbolJsObject.Get("math").Set("exp", mathquery.WasmQueryFunctionsStorage["WasmQueryExp"])
	symbolJsObject.Get("math").Set("exp2", mathquery.WasmQueryFunctionsStorage["WasmQueryExp2"])
	symbolJsObject.Get("math").Set("expm1", mathquery.WasmQueryFunctionsStorage["WasmQueryExpm1"])
	symbolJsObject.Get("math").Set("hypot", mathquery.WasmQueryFunctionsStorage["WasmQueryHypot"])
	symbolJsObject.Get("math").Set("ilogb", mathquery.WasmQueryFunctionsStorage["WasmQueryIlogb"])
	symbolJsObject.Get("math").Set("y0", mathquery.WasmQueryFunctionsStorage["WasmQueryY0"])
	symbolJsObject.Get("math").Set("y1", mathquery.WasmQueryFunctionsStorage["WasmQueryY1"])
	symbolJsObject.Get("math").Set("tanh", mathquery.WasmQueryFunctionsStorage["WasmQueryTanh"])
	symbolJsObject.Get("math").Set("trunc", mathquery.WasmQueryFunctionsStorage["WasmQueryTrunc"])
	symbolJsObject.Get("math").Set("tan", mathquery.WasmQueryFunctionsStorage["WasmQueryTan"])
	symbolJsObject.Get("math").Set("abs", mathquery.WasmQueryFunctionsStorage["WasmQueryAbs"])
	symbolJsObject.Get("math").Set("acos", mathquery.WasmQueryFunctionsStorage["WasmQueryAcos"])
	symbolJsObject.Get("math").Set("acosh", mathquery.WasmQueryFunctionsStorage["WasmQueryAcosh"])
	symbolJsObject.Get("math").Set("asin", mathquery.WasmQueryFunctionsStorage["WasmQueryAsin"])
	symbolJsObject.Set("geolocation", js.ValueOf(emptyObject))
	symbolJsObject.Get("geolocation").Set("latitude", geolocationquery.WasmQueryFunctionsStorage["WasmQueryLatitude"])
	symbolJsObject.Get("geolocation").Set("longitude", geolocationquery.WasmQueryFunctionsStorage["WasmQueryLongitude"])
	symbolJsObject.Set("secure", js.ValueOf(emptyObject))
	symbolJsObject.Get("secure").Set("secureLinks", securequery.WasmQueryFunctionsStorage["WasmQuerySecureLinks"])
	symbolJsObject.Get("secure").Set("redirectToHTTPS", securequery.WasmQueryFunctionsStorage["WasmQueryRedirectToHTTPS"])
	if len(args) > 0 {
		if args[0].Type() == js.TypeObject {
			if args[0].Length() > 1 {
				for i := 1; i < args[0].Length(); i++ {
					if args[0].Index(i).Type() == js.TypeString {
						js.Global().Set(args[0].Index(i).String(), symbolJsObject)
					}
				}
			}
		}
	}
	return js.Undefined()
}
