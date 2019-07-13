# WasmQuery
Experimental jQuery-like library written in Go. Available from JavaScript by WebAssembly module. Personally I recommend to use Vanilla JavaScript instead jQuery-like libraries. Vanilla JavaScript ES6 and higher are enough powerful.

## Syntax
General syntax  
```js
$("selector").method();
```
Standard invoke returns DOM elements  
For example:  
```js
$("h1");
$("#id");
$(".class");
```
To return CSS property
```js
$("h1")[0].css("color");
$("#id").css("color");
$(".class")[0].css("color");
```
To setup CSS property
```js
$("h1")[0].css("color", "fuchsia");
$("#id").css("color", "fuchsia");
$(".class")[0].css("color", "fuchsia");
$("h1").css("color", "fuchsia")
$(".class").css("color", "fuchsia")
```
To disable library and release memory
```js
$.disableLibrary();
```
