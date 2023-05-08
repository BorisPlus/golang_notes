# Интерфейсы. Двухсвязный список

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

```shell
go doc -all ./ > list.doc.txt
```

Документация:

```text
{{ list.doc.txt }}
```

{{ notice }}
