package pointer

import "fmt"

// Pointer - интерфейс точки произвольной размерности пространства.
type Pointer interface {
	Coordinates() interface{} // координаты точки.
	fmt.Stringer
}

// Eq(p1, p2 Pointer) - функция проверки равенства координат.
func Eq(p1, p2 Pointer) bool {
	return &p1 == &p2 || p1.Coordinates() == p2.Coordinates()
}
