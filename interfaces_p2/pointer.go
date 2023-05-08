package main

import "fmt"

// Решение задач классификации посредством интерфейса без привязки к метрике
// (без функции расчёта расстояния между точками и размерности пространства).

// Интерфейс точки (разной размерности и функции расстояния).
type Pointer interface {
	Coordinates() interface{}
	fmt.Stringer
}

func Eq(p1, p2 Pointer) bool {
	return p1.Coordinates() == p2.Coordinates()
}
