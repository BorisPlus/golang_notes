# Интерфейсы. Двусвязный список

## Реализация

Интерфейсы и структуры:

<details>
<summary>см. "list.go":</summary>

```go
{{ list.go }}
```

</details>

Наглядность:

<details>
<summary>см. "list_stringer.go":</summary>

```go
{{ list_stringer.go }}
```

</details>

Тестирование:

<details>
<summary>см. "list_test.go":</summary>

```go
{{ list_test.go }}
```

</details>

* TestListSimple

```shell
go test -v -run TestListSimple ./list.go ./list_stringer.go ./list_test.go  > list_test.go.simple.txt
```

<details>
<summary>см. лог:</summary>

```text
{{ list_test.go.txt }}
```

</details>

* TestListComplex

```shell
go test -v -run TestListComplex ./list.go ./list_stringer.go ./list_test.go  > list_test.go.complex.txt
```

<details>
<summary>см. лог:</summary>

```text
{{ list_test.go.complex.txt }}
```

</details>

* TestListComplex

```shell
go test -v -run TestListSwap ./list.go ./list_stringer.go ./list_test.go  > list_test.go.swap.txt
```

<details>
<summary>см. лог:</summary>

```text
{{ list_test.go.txt }}
```

</details>

## Документация

```shell
go doc -all ./ > list.doc.txt
```

<details>
<summary>см. документацию:</summary>

```text
{{ list.doc.txt }}
```

</details>

## Сортировка

Я, как мне кажется (😉), подобрал хороший пример для наглядной демонстрации интерфейса, требуемого sort.Sort (уже присутствуют в коде выше):

* `func (list *List) Less(i, j int) bool`
* `func (list *List) Swap(i, j int)`

В данном варианте не подойдет sort.Slice, так как перестановка элементов в двусвязном списке влечет перестановку указателей на соседей и с соседей на переставляемые элементы.

Особенность в том, что заранее зная, что будет реализован интерфейса sort.Sort, пришлось отказаться от именования Swap в самой структуре двусвязного списка, так как сигнатура должна быть `Swap (i, j int)`, а не как положено для двусвязного `Swap (i, j *ListItem)`. Это подводный камень для рефакторинга, стоит заранее избегать именований интерфейсных методов.

* TestListSortInterface

```shell
go test -v -run TestListSortInterface ./list.go ./list_stringer.go ./list_test.go  > list_test.go.sort.txt
```

<details>
<summary>см. лог (список упорядочился):</summary>

```text
{{ list_test.go.sort.txt }}
```

</details>

{{ notice }}
