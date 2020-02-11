# go-server

#### Funcionalidades - estrutura HTML:
- Seção com toda a tela (fundo)
  - id="map"
  - Google Maps no centro
  - 3 menus incorporados (apresentação na superfície superior) com tratamento de cor e de sombreado
- Menu superior fixo:
  - Logomarca
  - "Botão de Menu 1"
  - "Botão de Menu 2"
- Menu Lateral esquerdo:
  - Tópicos e subtópicos
- Menu lateral direito:
  - Vídeos
- Menu inferior:
  - Contato: Nomes e WhatsApp

```  
The W3C DOM standard is separated into 3 different parts:
Core DOM - standard model for all document types
XML DOM - standard model for XML documents
HTML DOM - standard model for HTML documents 
```

```   
The HTML DOM is a standard object model and programming interface for HTML.
It defines:
The HTML elements as objects
The properties of all HTML elements
The methods to access all HTML elements
The events for all HTML elements
In other words: The HTML DOM is a standard for how to get, change, add, or delete HTML elements.
DOM APIs are designed such that no script can ever detect the simultaneous execution of other scripts.
The security model of the web is based on the concept of "origins". Many of potential attacks on the web involve cross-origin actions.
HTML namespace: http://www.w3.org/1999/xhtml
```  

```  
DOM Interface:
The Document Object Model (DOM) is a representation - a model - of a document and its content. 
Dom trees have nodes: DocumentType node, Element nodes, Text nodes, Comment nodes, and ProcessingInstruction nodes. 
html tag    ->  interface HTMLHtmlElement : HTMLElement {};
Javascript  ->  document.getElementsByTagName(name) 
```  

```  
  Especificações também relevantes:
  - URL
  - HTTP
  - COOKIES
  - FETCH
  - WEBIDL
  - JAVASCRIPT / ECMA-262
  - UIEVENTS
  - DOM
  - TOUCH
  - FILEAPI
  - XHR
  - MQ
  - CSS
  - GEOMETRY
  - WEBGL
  - WSP (WebSocket protocol)
  - MATHML
  - SVG
```  

```  
IDL 
    [Exposed=Window]
    interface Example {
      // this is an IDl definition
    };
```  

```  
document.getElementById(id)
document.getElementsByTagName(name)
document.getElementsByClassName(name)
```  

```  
element.innerHTML =  new html content
element.setAttribute(attribute, value)
element.attribute = new value
```  

```  
document.createElement(element)	  Create an HTML element
document.removeChild(element)	  Remove an HTML element
document.appendChild(element)	  Add an HTML element
document.replaceChild(new, old)	  Replace an HTML element 
```  

```   
document.getElementById(id).onclick = function(){code}	Adding event handler code to an onclick event 
```  

```   
Animation:
function myMove() {
    var elem = document.getElementById("animate"); 
    var pos = 0;
    var id = setInterval(frame, 5);
    function frame() {
      if (pos == 350) {
        clearInterval(id);
      } else {
        pos++; 
        elem.style.top = pos + 'px'; 
        elem.style.left = pos + 'px'; 
      }
    }
  }
```  

```  
Tag events:
<h1 onclick="this.innerHTML = 'Ooops!'">Click on this text!</h1>
```  

```   
Event listeners
element.addEventListener(event, function, useCapture);
The first parameter is the type of the event (like "click" or "mousedown").
The second parameter is the function we want to call when the event occurs.
The third parameter is a boolean value specifying whether to use event bubbling or event capturing. This parameter is optional.
Ex.:
document.getElementById("myBtn").addEventListener("click", displayDate);
element.addEventListener("click", function(){ alert("Hello World!"); });
element.addEventListener("click", myFunction);
function myFunction() {
  alert ("Hello World!");
}
```  

```  
Navigation:
parentNode
childNodes[nodenumber]
firstChild
lastChild
nextSibling
previousSibling
Ex.: 
document.getElementById("id02").innerHTML = document.getElementById("id01").firstChild; 
```  

```  
Two big elements:
document.body - The body of the document
document.documentElement - The full document 
```  

```  
Get the name of a tag
document.getElementById("id02").innerHTML = document.getElementById("id01").nodeName;
```  

```  
NodeValue
nodeValue for element nodes is null
nodeValue for text nodes is the text itself
nodeValue for attribute nodes is the attribute value
```  

```  
TextNode
var node = document.createTextNode("This is a new paragraph.");
```  

```  
Remove child tag
<div id="div1">
    <p id="p1">This is a paragraph.</p>
    <p id="p2">This is another paragraph.</p>
  </div>
  
  <script>
  var parent = document.getElementById("div1");
  var child = document.getElementById("p1");
  parent.removeChild(child);
  </script>
```  

```  
Replace a child tag
<div id="div1">
    <p id="p1">This is a paragraph.</p>
    <p id="p2">This is another paragraph.</p>
  </div>
  
  <script>
  var para = document.createElement("p");
  var node = document.createTextNode("This is new.");
  para.appendChild(node);
  
  var parent = document.getElementById("div1");
  var child = document.getElementById("p1");
  parent.replaceChild(para, child);
  </script>
```  
  
```  
APIs
HTML Geolocation
HTML Drag and Drop
HTML Web Workers
HTML SSE
```  

```  
Creating new tags
<head>
...
<script>document.createElement("myTag")</script>
</head>
<body>
<myTag>My personal tag</myTag>
</body>
```  

```  
WASM
To source toolchains: source ./emsdk_env.sh --build=Release
Compiling:
We have to pass the linker flag -s WASM=1 to emcc (otherwise by default emcc will emit asm.js).
If we want Emscripten to generate an HTML page that runs our program, in addition to the Wasm binary and JavaScript wrapper, we have to specify an output filename with a .html extension.
Finally, to actually run the program, we cannot simply open the HTML file in a web browser because cross-origin requests are not supported for the file protocol scheme. We have to actually serve the output files over HTTP. 
```  