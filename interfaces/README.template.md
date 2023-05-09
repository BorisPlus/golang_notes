# Интерфейсы

## Варианты

Оба варианта

* `value receiver`

```go
{{ interface.go }}
```

* `pointed receiver` (код видоизменен в силу индикации линтера в IDE на однообразные именования, но суть понятна)

```go
{{ interface_pointed_receiver.go }}
```

приведут к успешной фильтрации

```text
{{ printed }}
```

## Поведение

```go
{{ interface_type_switch.go }}
```
