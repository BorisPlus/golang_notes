# Указатель

```go
package main

import (
    "math/rand"
)

func FunctionWithName() (func() int, *func() int) {
    anonymousFunction := func() int {
        return rand.Intn(10)
    }
    return anonymousFunction, &anonymousFunction
}

```

```go
fmt.Printf("%s",sFunctionWithName)
```

дает

```text
%!s(func() (func() int, *func() int)=0x493c60)
```

Оператор &NamedFunction недопустим, однако

```go
fmt.Printf("%p",pFunctionWithName)
```

```text
0x493c60 
```

0x493c60 - он всегда постоянный.

* что это такое? как это интерпритировать?
* если это адрес, то почему не 10 знаков после '0x'?\n", NamedFunction)  
* fmt.Printf("как это интерпритировать? если это адрес, то почему не 10 знаков после '0x'?
* fmt.Printf("если это адрес, то почему не 10 знаков после '0x'?
