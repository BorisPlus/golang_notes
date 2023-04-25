# Указатель

```go
content:./pointer.go
```

```go
fmt.Printf("%s",sFunctionWithName)
```

дает

```text
var:sFunctionWithName
```

Оператор &NamedFunction недопустим, однако

```go
fmt.Printf("%p",pFunctionWithName)
```

```text
var:pFunctionWithName 
```

var:pFunctionWithName :

* он всегда постоянный.
* что это такое? как это интерпритировать?
* если это адрес, то почему не 10 знаков после '0x'?
