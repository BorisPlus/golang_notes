# Интерфейсы. Двусвязный список

## Реализация

Интерфейс:

```go
{{ list.go }}
```

Наглядность:

```go
{{ list_stringer.go }}
```

Тестирование:

```go
{{ list_test.go }}
```

```shell
go test -v ./list.go ./list_stringer.go ./list_test.go  > list_test.go.txt
```

Лог:

```text
{{ list_test.go.txt }}
```

Документация:

```shell
go doc -all ./ > list.doc.txt
```

```text
{{ list.doc.txt }}
```

## Сортировка

Пример сортировки в соответствии с интерфейсом sort.Sort() - в случае, когда сортируем не слайсы и sort.Slice() не подходит.

Тестирование:

```go
{{ list_sort_test.go }}
```

```shell
go test -v ./list.go ./list_stringer.go ./list_sort_test.go  > list_sort_test.go.txt
```

Лог (список упорядочился):

```text
{{ list_sort_test.go.txt }}
```

{{ notice }}
