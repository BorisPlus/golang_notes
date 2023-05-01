# Интерфейсы. Классификация данных

Ниже представлен вариант реализации на Golang алгоритма соотнесения точки пространства с тем или иным классом точек, определяемым его центром.

Основной посыл, что классификатор заранее не знает размерность пространства и функцию подсчета расстояния, имеется только **интерфейс**.

```go
package main

import (
    "fmt"
    "math"
    "sort"
)

// Пробую решить задачу классификации без привязки к метрике
// (без функции расчёта расстояния между точками и размерности пространства).

// Интерфейс точки (разной размерности и функции расстояния).
type Pointer interface {
    fmt.Stringer
}

// Classificator - классификатор.
type Classificator struct {
    // centroids - это центры классов.
    centroids []Pointer
    // classes - это сами классы со всеми точками.
    classes map[Pointer][]Pointer
}

// Classificator.getNearestCentroid - получает ближайший центр класса в зависимости от
// переданной функции расстояния DistanceFunc (заранее не известной)
func (c *Classificator) getNearestCentroid(point Pointer, DistanceFunc func(a, b Pointer) float64) Pointer {
    distance := float64(-1)
    var pointCentroid Pointer
    for _, centroid := range c.centroids {
        pointToSampleDistance := DistanceFunc(centroid, point)
        if distance == -1 {
            distance = pointToSampleDistance
            pointCentroid = centroid
            continue
        }
        if pointToSampleDistance < distance {
            distance = pointToSampleDistance
            pointCentroid = centroid
        }
    }
    return pointCentroid
}

// Добавляет точку в тот или иной класс классификатора
// func (c *Classificator) classify(point Pointer, DistanceFunc func(a, b Pointer) float64) {
//     pointCentroid := c.getNearestCentroid(point, DistanceFunc)
//     c.classes[pointCentroid] = append(c.classes[pointCentroid], point)
// }

// ДВУМЕРНОЕ ПРОСТРАНСТВО ===========================

// XYPoint - пусть точка Point двумерная
type XYPoint struct {
    x, y float64
}

func (p XYPoint) String() string {
    return fmt.Sprintf("[x: %v, y: %v]", p.x, p.y)
}

// EuclidianXYPow2 - пусть расстояние между точками это Евклидова-метрика без ее корня
func EuclidianXYPow2(p1, p2 Pointer) float64 {
    return ((p1.(XYPoint).x-p2.(XYPoint).x)*(p1.(XYPoint).x-p2.(XYPoint).x) +
        (p1.(XYPoint).y-p2.(XYPoint).y)*(p1.(XYPoint).y-p2.(XYPoint).y))
}

// ТРЕХМЕРНОЕ ПРОСТРАНСТВО ===========================

// XYZPoint - пусть точка Point трехмерная
type XYZPoint struct {
    x, y, z float64
}

func (p XYZPoint) String() string {
    return fmt.Sprintf("[x: %v, y: %v, z: %v]", p.x, p.y, p.z)
}

// Decart - пусть расстояние между точками - это вариация на Декартову метрику
func Decart(p1, p2 Pointer) float64 {
    dXdYdZ := []float64{
        math.Abs(p1.(XYZPoint).x - p2.(XYZPoint).x),
        math.Abs(p1.(XYZPoint).y - p2.(XYZPoint).y),
        math.Abs(p1.(XYZPoint).z - p2.(XYZPoint).z)}
    sort.Float64s(dXdYdZ)
    return dXdYdZ[0]
}

// ПРИМЕР ===========================

func main() {
    // Это координаты вершин квадрата 10x10
    centroids := []Pointer{
        XYPoint{x: 0, y: 0},
        XYPoint{x: 10, y: 0},
        XYPoint{x: 0, y: 10},
        XYPoint{x: 10, y: 10}}
    classificator := Classificator{centroids, make(map[Pointer][]Pointer, 0)}
    fmt.Println(classificator)
    // Берем точку с координатами
    point_1_1 := Pointer(XYPoint{x: 1, y: 1})
    // Выясняем какой центройд ей ближе по Евклиду
    fmt.Println(classificator.getNearestCentroid(point_1_1, EuclidianXYPow2))
    // А теперь трехмерный вариант
    // Это координаты вершин квадрата 10x10
    centroids = []Pointer{
        XYZPoint{x: 0, y: 0, z: 0},
        XYZPoint{x: 10, y: 0, z: 10},
        XYZPoint{x: 0, y: 10, z: 5},
        XYZPoint{x: 10, y: 10, z: 10}}
    classificator = Classificator{centroids, make(map[Pointer][]Pointer, 0)}
    fmt.Println(classificator)
    // Берем точку с координатами
    point_1_1_2 := Pointer(XYZPoint{x: 1, y: 1, z: 2})
    // Выясняем какой центроид ей ближе по Евклиду
    fmt.Println(classificator.getNearestCentroid(point_1_1_2, Decart))
    // А теперь ... вводим понятие расстояния между "синим" и "теплым" Ж)
}

```

```go
> ```text
> Данный документ составлен с использованием 
> разработанного [шаблонизатора](https://github.com/BorisPlus/golang_notes/tree/master/templator). 
> При его использовании избегайте рекурсивной вложенности.
> ```
```
