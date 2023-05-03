# Интерфейсы. Классификация данных

Ниже представлен вариант реализации на Golang алгоритма соотнесения точки пространства с тем или иным классом точек, определяемым его центром.

Основной посыл, что классификатор заранее не знает размерность пространства и функцию подсчета расстояния, имеется только **интерфейс**.

Решение задач классификации посредством интерфейса без привязки к метрике (без функции расчёта расстояния между точками и размерности пространства).

## Реализация

Пояснения представлены в рамках комментариев.

Интерфейс точки:

```go
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

```

Реализация классификатора:

```go
package main

// Classificator - классификатор.
type Classificator struct {
    // centroids - это центры классов.
    centroids []Pointer
}

// Classificator.getNearestCentroid - получает ближайший центр класса в зависимости от
// переданной функции расстояния - DistanceBetween (заранее не известной)
func (c *Classificator) getNearestCentroid(point Pointer, DistanceBetween func(a, b Pointer) float64) Pointer {
    distance := float64(-1)
    var pointCentroid Pointer
    for _, centroid := range c.centroids {
        pointToSampleDistance := DistanceBetween(centroid, point)
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

```

Решение задач классификации посредством интерфейса без привязки к метрике (тестирование):

```go
package main

import (
    "fmt"
    "math"
    "testing"
)

// ДВУМЕРНОЕ ПРОСТРАНСТВО ===========================

// XYPoint - пусть точка Point двумерная
type XYPoint struct {
    x, y float64
}

func (p XYPoint) String() string {
    return fmt.Sprintf("(%v;%v)", p.x, p.y)
}

func (p XYPoint) getValue() string {
    return p.String()
}

// EuclidianXYPow2 - пусть расстояние между точками это Евклидова-метрика без ее корня
func EuclidianXYPow2(p1, p2 Pointer) float64 {
    return ((p1.(XYPoint).x-p2.(XYPoint).x)*(p1.(XYPoint).x-p2.(XYPoint).x) +
        (p1.(XYPoint).y-p2.(XYPoint).y)*(p1.(XYPoint).y-p2.(XYPoint).y))
}

func TestXY(t *testing.T) {
    // Это координаты вершин квадрата 10x10:
    //
    // ↑ y
    // |
    // |(0;10)
    // ├------┐ (10;10)
    // |      |
    // |      |     x
    // └------┚-----→
    // (0;0)   (10;0)
    //
    awaitedCentroid := Pointer(XYPoint{x: 0, y: 0})
    fmt.Println("AwaitedCentroid", awaitedCentroid)
    centroids := []Pointer{
        XYPoint{x: 10, y: 0},
        XYPoint{x: 0, y: 10},
        XYPoint{x: 10, y: 10}}
    //
    centroids = append(centroids, awaitedCentroid)
    fmt.Println("Centroids", centroids)
    classificator := Classificator{centroids, make(map[Pointer][]Pointer, 0)}
    fmt.Println("Classificator", classificator)
    // Берем точку с координатами
    point := Pointer(XYPoint{x: 1, y: 1})
    fmt.Println("Point", point)
    // Выясняем какой центройд ей ближе по Евклиду
    classifiedTo := classificator.getNearestCentroid(point, EuclidianXYPow2)
    fmt.Println("ClassifiedTo", classifiedTo)
    if !Eq(classifiedTo, awaitedCentroid) {
        t.Errorf("Point %v must be classified to %v, but was %v. Error.", point, awaitedCentroid, classifiedTo)
    } else {
        fmt.Printf("Point %v classified to  %v (%v). OK.\n", point, classifiedTo, awaitedCentroid)
    }
}

// ТРЕХМЕРНОЕ ПРОСТРАНСТВО ===========================

// XYZPoint - пусть точка Point трехмерная
type XYZPoint struct {
    x, y, z float64
}

func (p XYZPoint) String() string {
    return fmt.Sprintf("(%v;%v;%v)", p.x, p.y, p.z)
}

func (p XYZPoint) getValue() string {
    return p.String()
}

// Decart - пусть расстояние между точками - это вариация на Декартову метрику
func Decart(p1, p2 Pointer) float64 {
    dXdYdZ := math.Abs(p1.(XYZPoint).x-p2.(XYZPoint).x) +
        math.Abs(p1.(XYZPoint).y-p2.(XYZPoint).y) +
        math.Abs(p1.(XYZPoint).z-p2.(XYZPoint).z)
    return dXdYdZ
}

func TestXYZ(t *testing.T) {
    awaitedCentroid := Pointer(XYZPoint{x: 0, y: 0, z: 0})
    centroids := []Pointer{
        XYZPoint{x: 10, y: 0, z: 10},
        XYZPoint{x: 0, y: 10, z: 5},
        XYZPoint{x: 10, y: 10, z: 10}}
    centroids = append(centroids, awaitedCentroid)
    fmt.Println("Centroids", centroids)
    //
    classificator := Classificator{centroids, make(map[Pointer][]Pointer, 0)}
    fmt.Println("Classificator", classificator)
    point := Pointer(XYZPoint{x: 0, y: 0, z: 0})
    fmt.Println("Point", point)
    classifiedTo := classificator.getNearestCentroid(point, Decart)
    fmt.Println("ClassifiedTo", classifiedTo)
    // if classifiedTo != awaitedCentroid {
    if !Eq(classifiedTo, awaitedCentroid) {
        t.Errorf("Point %v must be classified to  %v, but was %v. ERROR.", point, awaitedCentroid, classifiedTo)
    } else {
        fmt.Printf("Point %v classified to  %v (%v). OK.\n", point, classifiedTo, awaitedCentroid)
    }
}

```

```shell
go test -v ./pointer.go ./classificator.go ./classificator_test.go  > ./classificator.go.txt
```

Лог:

```text
=== RUN   TestXY
AwaitedCentroid (0;0)
Centroids [(10;0) (0;10) (10;10) (0;0)]
Classificator {[{10 0} {0 10} {10 10} {0 0}] map[]}
Point (1;1)
ClassifiedTo (0;0)
Point (1;1) classified to  (0;0) ((0;0)). OK.
--- PASS: TestXY (0.00s)
=== RUN   TestXYZ
Centroids [(10;0;10) (0;10;5) (10;10;10) (0;0;0)]
Classificator {[{10 0 10} {0 10 5} {10 10 10} {0 0 0}] map[]}
Point (0;0;0)
ClassifiedTo (0;0;0)
Point (0;0;0) classified to  (0;0;0) ((0;0;0)). OK.
--- PASS: TestXYZ (0.00s)
PASS
ok  	command-line-arguments	0.005s

```

## Вывод

Интерфейс как составляющая часть инкапсуляции добавляет гибкости для будущих решений.

Приведенные выше измерения пространства могут быть перенесены на сравнения категорий, шутливо - "синего" и "тёплого", необходимо только придерживаться сигнатуры функции:

```go
... DistanceBetween func(a, b Pointer) float64 ... 
```

## Послесловие

>
> ```text
> Данный документ составлен с использованием разработанного шаблонизатора. 
> При его использовании избегайте рекурсивной вложенности.
> ```

см. ["Шаблонизатор"](https://github.com/BorisPlus/golang_notes/tree/master/templator)
