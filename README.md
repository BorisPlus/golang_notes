# Стек

В рамках вебинара аудиторией независимо разработаны различные варианты реализации стека на Go.

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

## Тестирвание

Тестирование происходит по одной схеме

```go
import (
    "testing"
)

func BenchmarkSimplest(b *testing.B) {
    stack := Stack{}
    // заполняем стек
    for _, num := range array {
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
BenchmarkSimplestKonstantin10Values-4           1000000000               0.0000095 ns/op
BenchmarkSimplestKonstantin1000Values
BenchmarkSimplestKonstantin1000Values-4         1000000000               0.001856 ns/op
BenchmarkSimplestKonstantin100000Values
BenchmarkSimplestKonstantin100000Values-4              1        16548906905 ns/op
BenchmarkSimplestRoman10Values
BenchmarkSimplestRoman10Values-4                1000000000               0.0000052 ns/op
BenchmarkSimplestRoman1000Values
BenchmarkSimplestRoman1000Values-4              1000000000               0.0000362 ns/op
BenchmarkSimplestRoman100000Values
BenchmarkSimplestRoman100000Values-4            1000000000               0.002220 ns/op
PASS
ok      stack/stack     16.625s
```

Наглядно таблицей

| Dataset \ ns/op   | Konstantin    |   Roman   |
| :---------------- | :-----------: | --------: |
| 10                | 0.0000095     | 0.0000052 |
| 1000              | 0.001856      | 0.0000362 |
| 100000            | 16548906905   | 0.002220  |

## Вывод

Вариант от Roman на значительном объеме эффективнее.

Roman, мои поздравления 🏆.

## Дополнительно

Если избавиться от stack.len, введя defer в Pop

```go
package main

type StackRomanWL struct {
    items []int
}

func (stack *StackRomanWL) Push(i int) {
    stack.items = append(stack.items, i)
}

func (stack *StackRomanWL) Pop() (filo int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = r.(error)
        }
    }()
    stack_items_len := len(stack.items)
    filo = stack.items[stack_items_len-1]
    stack.items = stack.items[:stack_items_len-1]
    return
}
```

то это замедлит

```shell
go test -v -bench=. stack_roman.go stack_roman_without_len.go \
stack_roman_test.go stack_roman_without_len_test.go \
dataset.go
```

```text
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i3-2310M CPU @ 2.10GHz
BenchmarkSimplestRoman10Values
BenchmarkSimplestRoman10Values-4                        1000000000               0.0000053 ns/op
BenchmarkSimplestRoman1000Values
BenchmarkSimplestRoman1000Values-4                      1000000000               0.0000205 ns/op
BenchmarkSimplestRoman100000Values
BenchmarkSimplestRoman100000Values-4                    1000000000               0.002607 ns/op
BenchmarkSimplestRomanWithoutLen10Values
BenchmarkSimplestRomanWithoutLen10Values-4              1000000000               0.0000119 ns/op
BenchmarkSimplestRomanWithoutLen1000Values
BenchmarkSimplestRomanWithoutLen1000Values-4            1000000000               0.0000954 ns/op
BenchmarkSimplestRomanWithoutLen100000Values
BenchmarkSimplestRomanWithoutLen100000Values-4          1000000000               0.003979 ns/op
PASS
ok      command-line-arguments  0.093s
```

Изящность кода с `defer` и `"naked" return` не добавляет скорости.
