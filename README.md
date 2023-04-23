# Стек

В рамках вебинара аудиторией независимо разработаны различные варианты реализации стека на Go.

В целях установления эффективности того или иного подхода возникла идея замера производительности.

__Замечания к изменениям алгоритмов__:

* именования переменных и структур приведены к единообразию;
* для последующей проверки доступности элементов на извлечение введено поле длины очереди.

__Замечание__: ваши правки приветствуются.

## Базовые варианты реализации

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

### Тестирвание базовых вариантов

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

### Вывод в рамках базовых вариантов

Вариант от Roman на значительном объеме эффективнее.

Roman, мои поздравления 🏆.

## Вариации на тему

Теперь рассмотрим попытки повысить скорость работы "победившей" реализации.

### Нуль указатель вершины (и что все таки быстрее?)

В изначальных вариантах для проверки доступности элементов на извлечение было введено поле длины очереди `len++/--`.
Заменим этот критерий на возврат `nil`, если очередь пуста.

```go
package main

type StackPopedNil struct {
    items []int
    // Больше нет поля `len`
}

func (stack *StackPopedNil) Push(value int) {
    stack.items = append(stack.items, value)
}

func (stack *StackPopedNil) Pop() *int { // Изменился возвращаемый тип
    stack_items_len := len(stack.items)
    if stack_items_len == 0 {
        return nil // Возвращает nil указатель
    }
    filo := stack.items[stack_items_len-1]
    stack.items = stack.items[:stack_items_len-1]
    return &filo // Возвращает указатель на элемент
}
```

И теперь __самое__ инетерсное:

```bash
go test -v -bench=. dataset.go stack_poped_nil.go stack_poped_nil_test.go stack_roman.go stack_roman_10000000_test.go
```

Два запуска противоречат друг другу

```text
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i3-2310M CPU @ 2.10GHz
BenchmarkStackPopedNil10000000Values
BenchmarkStackPopedNil10000000Values-4          1000000000               0.2489 ns/op <---┐
BenchmarkRoman10000000Values                                                              |
BenchmarkRoman10000000Values-4                  1000000000               0.2983 ns/op <---|-----┐
PASS                                                                                      |     |
ok      command-line-arguments  8.300s                                                    |     |
                                                                                          |     |
                                                                                          |     |
stack_roman_10000000_test.go                                                              |     |  Это
goos: linux                                                                               |     |  диаметрально
goarch: amd64                                                                             |     |  противоположный
cpu: Intel(R) Core(TM) i3-2310M CPU @ 2.10GHz                                             |     |  результат.
BenchmarkStackPopedNil10000000Values                                                      |     |  Что быстрее?
BenchmarkStackPopedNil10000000Values-4          1000000000               0.2902 ns/op <---┘     |
BenchmarkRoman10000000Values                                                                    |
BenchmarkRoman10000000Values-4                  1000000000               0.2235 ns/op <---------┘
PASS
ok      command-line-arguments  8.082s
```

### Иные варианты

#### Pointered - хранение указателей

```go
package main

type StackPointered struct {
    items []*int // Учет указателей
    len uint
}

func (stack *StackPointered) Push(value *int) { // Учет указателей
    stack.items = append(stack.items, value)
    stack.len++
}

func (stack *StackPointered) Pop() *int { // Возврат указателей
    stack_items_len := len(stack.items)
    filo := stack.items[stack_items_len-1]
    stack.items = stack.items[:stack_items_len-1]
    stack.len--
    return filo // Возврат указателя
}
```

```shell
go test -v -bench=. dataset.go stack_pointered.go stack_pointered_test.go stack_roman.go stack_roman_10000000_test.go
```

```text
=== RUN   TestStackPointered10Values
--- PASS: TestStackPointered10Values (0.00s)
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i3-2310M CPU @ 2.10GHz
BenchmarkStackPointered10000000Values
BenchmarkStackPointered10000000Values-4                1        1253531463 ns/op
BenchmarkRoman10000000Values
BenchmarkRoman10000000Values-4                  1000000000               0.2342 ns/op
PASS
ok      command-line-arguments  6.034s
```

#### Defer - как отслеживание INDEX OUT OF RANGE

```go
func (stack *Stack) Pop() (filo int, err error) {
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

Успеха производительности не будет

#### Special structed

Идея в хранении в стеке указателя на текущую FI-вершину, содержащую ссылку на следующую по стеку

```go
package main

type Item struct {
    value int
    next  *Item
}

type StackStructed struct {
    origin *Item
}

func (stack *StackStructed) Push(value int) {
    willBePushed := Item{value: value, next: stack.origin}
    stack.origin = &willBePushed
}

func (stack *StackStructed) Pop() (int, *Item) {
    poped := stack.origin.value
    stack.origin = stack.origin.next
    return poped, stack.origin 
}
```

Результат тестирования

```shell
go test -v -bench=. dataset.go stack_structed.go stack_structed_test.go stack_roman.go stack_roman_10000000_test.go
```

работает медленнее:

```text
go
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i3-2310M CPU @ 2.10GHz
BenchmarkStackStructed10000000Values
BenchmarkStackStructed10000000Values-4                 4         272123579 ns/op
BenchmarkRoman10000000Values
BenchmarkRoman10000000Values-4                1000000000         0.2544 ns/op
PASS
ok      command-line-arguments  8.308s
```

## Вывод

Изначально, в производительности "состязались" варианты Константина и Романа, которые были немного унифицированы (см. ["Замечания к изменениям алгоритмов"](#%D1%81%D1%82%D0%B5%D0%BA)).

Далее проводилась модификация победившего среди базовых алгоритма Романа в части касающейся критерия останова.

Продемонстрировано, что незначительные изменения в коде ведут к поразительному снижению производительности.

Таким образом, изначальное введение дополнительного поля `stack.len(++/--)` - это весьма и весьма удачно принятое решение.

Остается только вопрос в отношении нестабильных показателей "оценки состязательности" алгоритмов главы ["Нуль указатель вершины (и что все таки быстрее?)"](#%D0%BD%D1%83%D0%BB%D1%8C-%D1%83%D0%BA%D0%B0%D0%B7%D0%B0%D1%82%D0%B5%D0%BB%D1%8C-%D0%B2%D0%B5%D1%80%D1%88%D0%B8%D0%BD%D1%8B-%D0%B8-%D1%87%D1%82%D0%BE-%D0%B2%D1%81%D0%B5-%D1%82%D0%B0%D0%BA%D0%B8-%D0%B1%D1%8B%D1%81%D1%82%D1%80%D0%B5%D0%B5). На этом пока все.

Всем, оислившим чтиво, спасибо.
