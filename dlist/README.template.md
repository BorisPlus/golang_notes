# Интерфейсы. Двусвязный список

## Реализация

Интерфейсы и структуры:

<details>
<summary>см. "dlist.go":</summary>

```go
{{ dlist.go }}
```

</details>

Наглядность:

<details>
<summary>см. "dlist_stringer.go":</summary>

```go
{{ dlist_stringer.go }}
```

</details>

Тестирование:

<details>
<summary>см. "dlist_test.go":</summary>

```go
{{ dlist_test.go }}
```

</details>

* TestDListSimple

```shell
go test -v -run TestDListSimple > dlist_test.go.simple.txt
```

<details>
<summary>см. лог "TestDListSimple":</summary>

```text
{{ dlist_test.go.simple.txt }}
```

</details>

* TestDListComplex

```shell
go test -v -run TestDListComplex > dlist_test.go.complex.txt
```

<details>
<summary>см. лог "TestDListComplex":</summary>

```text
{{ dlist_test.go.complex.txt }}
```

</details>

* TestDListComplex

```shell
go test -v -run TestDListSwap > dlist_test.go.swap.txt
```

<details>
<summary>см. лог "TestDListSwap":</summary>

```text
{{ dlist_test.go.swap.txt }}
```

</details>

## Документация

```shell
go doc -all ./ > dlist.doc.txt
```

<details>
<summary>см. документацию:</summary>

```text
{{ dlist.doc.txt }}
```

</details>

## Сортировка

* TestListSortInterface
  
Я, как мне кажется (😉), подобрал хороший пример для наглядной демонстрации интерфейса, требуемого sort.Sort (уже присутствуют в коде выше):

* `func (list *DList) Len()`
* `func (list *DList) Less(i, j int) bool`
* `func (list *DList) Swap(i, j int)`

В данном варианте не подойдет sort.Slice, так как перестановка элементов в двусвязном списке влечет перестановку указателей на соседей и с соседей на переставляемые элементы.

Особенность реализации:

* заранее зная, что будет реализован интерфейс sort.Sort, пришлось отказаться от именования Swap в самой структуре двусвязного списка, так как сигнатура должна быть `Swap (i, j int)`, а не как положено для двусвязного `Swap (i, j *DListItem)`. Это подводный камень для рефакторинга - стоит заранее избегать именований интерфейсных методов;
* поскольку элементы могут хранить различные структуры, а для сортировки необходимо проводить сравнения `Less(i, j int)`, в структуру списка введена функция по получению результата сравнения элементов списка `less func(x, y *DListItem) bool`. Так в тесте приведены варианты упорядочивания двусвязного списка со значениями с `int`, `rune` и `custom-stuct` без необходимости вмещательства в целевой пакет.

```shell
go test -v -run TestDListSortInterface > dlist_test.go.sort.txt
```

<details>
<summary>см. лог "TestDListSortInterface" (списки упорядочились):</summary>

```text
{{ dlist_test.go.sort.txt }}
```

</details>

{{ notice }}
