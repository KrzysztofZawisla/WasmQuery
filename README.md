# WasmQuery
Experimental jQuery-like library written in Go. Available from JavaScript by WebAssembly module. Personally I recommend to use Vanilla JavaScript instead jQuery-like libraries. Vanilla JavaScript ES6 and higher are powerful enough.

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
$("h1").css("color");
$(".class").css("color");
```
To setup CSS property
```js
$("h1")[0].css("color", "fuchsia");
$("#id").css("color", "fuchsia");
$(".class")[0].css("color", "fuchsia");
$("h1").css("color", "fuchsia")
$(".class").css("color", "fuchsia")
```
To hide DOM elements
```js
$("h1")[0].hide();
$("#id").hide();
$(".class")[0].hide();
$("h1").hide();
$(".class").hide();
```
To show DOM elements as block 
```js
$("h1")[0].show();
$("#id").show();
$(".class")[0].show();
$("h1").show();
$(".class").show();
```
or
```js
$("h1")[0].showAsBlock();
$("#id").showAsBlock();
$(".class")[0].showAsBlock();
$("h1").showAsBlock();
$(".class").showAsBlock();
```
To show DOM elements as inline 
```js
$("h1")[0].showAsInline();
$("#id").showAsInline();
$(".class")[0].showAsInline();
$("h1").showAsInline();
$(".class").showAsInline();
```
To show DOM elements as inline-block
```js
$("h1")[0].showAsInlineBlock();
$("#id").showAsInlineBlock();
$(".class")[0].showAsInlineBlock();
$("h1").showAsInlineBlock();
$(".class").showAsInlineBlock();
```
To show DOM elements as flex
```js
$("h1")[0].showAFlex();
$("#id").showAFlex();
$(".class")[0].showAFlex();
$("h1").showAFlex();
$(".class").showAFlex();
```
To disable library and release memory
```js
$.disableLibrary();
```
