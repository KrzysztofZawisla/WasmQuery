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
$("h1")[0];
$("#id")[0];
$(".class")[0];
$("h1");
$("#id");
$(".class");
```
To return CSS property
```js
$("h1")[0].css("color");
$("#id")[0].css("color");
$(".class")[0].css("color");
$("h1").css("color");
$("#id").css("color");
$(".class").css("color");
```
To setup CSS property
```js
$("h1")[0].css("color", "fuchsia");
$("#id")[0].css("color", "fuchsia");
$(".class")[0].css("color", "fuchsia");
$("h1").css("color", "fuchsia")
$("#id").css("color", "fuchsia");
$(".class").css("color", "fuchsia")
```
or
```js
$("h1")[0].css("color", ["fuchsia"])
$("#id")[0].css("color", ["fuchsia"])
$(".class")[0].css("color", ["fuchsia"])
$("h1").css("color", ["fuchsia", "fuchsia", "fuchsia"]);
$("#id").css("color", ["fuchsia"])
$(".class").css("color", ["fuchsia", "fuchsia", "fuchsia"]);
```
To hide DOM elements
```js
$("h1")[0].hide();
$("#id")[0].hide();
$(".class")[0].hide();
$("h1").hide();
$("#id").hide();
$(".class").hide();
```
To show DOM elements based on tag name
```js
$("h1")[0].show();
$("#id")[0].show();
$(".class")[0].show();
$("h1").show();
$("#id").show();
$(".class").show();
```
To show DOM elements as block 
```js
$("h1")[0].showAsBlock();
$("#id")[0].showAsBlock();
$(".class")[0].showAsBlock();
$("h1").showAsBlock();
$("#id").showAsBlock();
$(".class").showAsBlock();
```
To show DOM elements as inline 
```js
$("h1")[0].showAsInline();
$("#id")[0].showAsInline();
$(".class")[0].showAsInline();
$("h1").showAsInline();
$("#id").showAsInline();
$(".class").showAsInline();
```
To show DOM elements as inline-block
```js
$("h1")[0].showAsInlineBlock();
$("#id")[0].showAsInlineBlock();
$(".class")[0].showAsInlineBlock();
$("h1").showAsInlineBlock();
$("#id").showAsInlineBlock();
$(".class").showAsInlineBlock();
```
To show DOM elements as flex
```js
$("h1")[0].showAsFlex();
$("#id")[0].showAsFlex();
$(".class")[0].showAsFlex();
$("h1").showAsFlex();
$("#id").showAsFlex();
$(".class").showAsFlex();
```
To return value
```js
$("h1")[0].val();
$("#id")[0].val();
$(".class")[0].val();
$("h1").val();
$("#id").val();
$(".class").val();
```
To setup value
```js
$("h1")[0].val("exampleValue");
$("#id")[0].val("exampleValue");
$(".class")[0].val("exampleValue");
$("h1").val("exampleValue");
$("#id").val("exampleValue");
$(".class").val("exampleValue");
```
or
```js
$("h1")[0].val(["exampleValue"]);
$("#id")[0].val(["exampleValue"]);
$(".class")[0].val(["exampleValue"]);
$("h1").val(["exampleValue1", "exampleValue2", "exampleValue3"]);
$("#id").val(["exampleValue1"]);
$(".class").val(["exampleValue1", "exampleValue2", "exampleValue3"]);
```
To return lenght of DOM elements array
```js
$("h1").len();
$("#id").len();
$(".class").len();
```
To return width of DOM elements
```js
$("h1")[0].width();
$("#id")[0].width();
$(".class")[0].width();
$("h1").width();
$("#id").width();
$(".class").width();
```
To setup width of DOM elements
```js
$("h1")[0].width("100px");
$("#id")[0].width("100px");
$(".class")[0].width("100px");
$("h1").width("100px");
$("#id").width("100px");
$(".class").width("100px");
```
or
```js
$("h1")[0].width(["100px"]);
$("#id")[0].width(["100px"]);
$(".class")[0].width(["100px"]);
$("h1").width(["100px", "200px", "300px"]);
$("#id").width(["100px"]);
$(".class").width(["100px", "200px", "300px"]);
```
To disable library and release memory
```js
$.disableLibrary();
```
