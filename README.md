# Стек

В рамках вебинара аудиторией независимо реализованы различные варианты реализации стека на Go.

В целях установления эффективности того или иного подхода возникла идея замера производительности.

__Замечание__: именования переменных и структур приведены к примерному единообразию, дополниительно везде введено поле длины очереди для последующей проверки доступности на извлечение.

__Замечание__: замечания приветствуются.

## Варианты реализации

### Konstantin

[https://go.dev/play/p/j2AVLXLh4Tt](https://go.dev/play/p/j2AVLXLh4Tt)

```go
package main

type StackKonstantin struct {
    items []int
    len   uint
}

func (stack *StackKonstantin) Push(i int) {
    stack.items = append([]int{i}, stack.items...)
    stack.len++
}

func (stack *StackKonstantin) Pop() int {
    filo := stack.items[0]
    stack.items = stack.items[1:]
    stack.len--
    return filo
}
```

### Roman

[https://go.dev/play/p/_ZwCVgiT4Dt](https://go.dev/play/p/_ZwCVgiT4Dt)

```go
package main

type StackRoman struct {
    items []int
    len   uint
}

func (stack *StackRoman) Push(i int) {
    stack.items = append(stack.items, i)
    stack.len++
}

func (stack *StackRoman) Pop() int {
    filo := stack.items[len(stack.items)-1]
    stack.items = stack.items[:len(stack.items)-1]
    stack.len--
    return filo
}
```

### Mine

```go
package main

type StackMine struct {
    items []int
    len uint
}

func (stack *StackMine) Push(i int) {
    stack.items = append(stack.items, i)
    stack.len++
}

func (stack *StackMine) Pop() int {
    filo := stack.items[0]
    stack.items = stack.items[1:]
    stack.len--
    return filo
}
```

## Тестирвание

Тестирование происходит по одной схеме

```go
import (
    "testing"
)

func BenchmarkSimplest(b *testing.B) {
    stack := Stack{}
    // заполняем стек
    for num := range array {
        stack.Push(num)
    }
    // высвобождаем стек
    for stack.len > 0 {
        stack.Pop()
    }
}
```

и на одинаковых данных

```go
package main

import (
    "math/rand"
)

var array10 = rand.Perm(10)
var array1000 = rand.Perm(1000)
var array100000 = rand.Perm(100000)
```

Итак:

```shell
golangci-lint run --out-format=github-actions
go test -v -bench=.
```

момент "истины":

```text
goos: linux
goarch: amd64
pkg: stack/stack
cpu: Intel(R) Core(TM) i3-2310M CPU @ 2.10GHz
BenchmarkSimplestKonstantin10Values
BenchmarkSimplestKonstantin10Values-4           1000000000               0.0000074 ns/op
BenchmarkSimplestKonstantin1000Values
BenchmarkSimplestKonstantin1000Values-4         1000000000               0.001418 ns/op
BenchmarkSimplestKonstantin100000Values
BenchmarkSimplestKonstantin100000Values-4              1        17876335812 ns/op
BenchmarkSimplestMine10Values
BenchmarkSimplestMine10Values-4                 1000000000               0.0000044 ns/op
BenchmarkSimplestMine1000Values
BenchmarkSimplestMine1000Values-4               1000000000               0.0000270 ns/op
BenchmarkSimplestMine100000Values
BenchmarkSimplestMine100000Values-4             1000000000               0.003949 ns/op
BenchmarkSimplestRoman10Values
BenchmarkSimplestRoman10Values-4                1000000000               0.0000075 ns/op
BenchmarkSimplestRoman1000Values
BenchmarkSimplestRoman1000Values-4              1000000000               0.0000265 ns/op
BenchmarkSimplestRoman100000Values
BenchmarkSimplestRoman100000Values-4            1000000000               0.003324 ns/op
PASS
ok      stack/stack     17.987s
```

Наглядно таблицей

| Dataset \ ns/op   | Konstantin    | Mine      |   Roman   |
| :---------------- | :-----------: | :-------: | --------: |
| 10                | 0.0000074     | 0.0000044 | 0.0000075 |
| 1000              | 0.001418      | 0.0000270 | 0.0000265 |
| 100000            | 17876335812   | 0.003949  | 0.003324  |

## Вывод

Вариант от Roman на значительном объеме эффективнее остальных.
Roman, мои поздравления 🏆.
