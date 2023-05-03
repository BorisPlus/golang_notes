package main

import "fmt"

// Решение задач классификации посредством интерфейса без привязки к метрике
// (без функции расчёта расстояния между точками и размерности пространства).

type Valuer interface {
	getValue() string
}

// Интерфейс точки (разной размерности и функции расстояния).
type Pointer interface {
	Valuer
	fmt.Stringer
}

func Eq(p1, p2 Pointer) bool {
	return p1.getValue() == p2.getValue()
}
