# Указатель

Исполнение

```go
{{ pointer.go }}
```

```go
fmt.Printf("%s",sFunctionWithName)
```

дает

```text
{{ sFunctionWithName }}
```

Оператор & для NamedFunction недопустим, однако

```go
fmt.Printf("%p",pFunctionWithName)
```

```text
{{ pFunctionWithName }}
```

{{ pFunctionWithName }} :

* он всегда постоянный.
* что это такое? как это интерпритировать?
* если это адрес, то почему не 10 знаков после '0x'?
