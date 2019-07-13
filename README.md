# WasmQuery
Experimental jQuery like library written in Go. Available from JavaScript by WebAssembly module.

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
To return CSS propertie
```js
$("h1")[0].css("color");
$("#id").css("color");
$(".class")[0].css("color");
```
To setup CSS propertie
```js
$("h1")[0].css("color", "fuchsia");
$("#id").css("color", "fuchsia");
$(".class")[0].css("color", "fuchsia");
```
To disable library and release memory
```js
$.disableLibrary();
```
