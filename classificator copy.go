package main

import (
	"fmt"
)

// Пробую решить задачу классификации без привязки к метрике.

// Интерфейс точки (разной размерности и функции расстояния).
type Pointer interface {
}

// Classificator - классификатор.
type Classificator struct {
	// centroids - это центры классов.
	centroids []Pointer
	// classes - это сами классы со всеми точками.
	classes   map[Pointer][]Pointer
}

// Classificator.getNearestCentroid - получает ближайший центр класса в зависимости от 
// переданной функции расстояния DistanceFunc (заранее не известной)
func (c *Classificator) getNearestCentroid(point Pointer, DistanceFunc func(a, b Pointer) int) Pointer {
	distance := -1
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
func (c *Classificator) classify(point Pointer, DistanceFunc func(a, b Point) int) {
	pointCentroid := c.getNearestCentroid(point, DistanceFunc)
	c.classes[pointCentroid] = append(c.classes[pointCentroid], point)
}

// XYPoint - пусть точка Point двумерная
// type XYPoint struct {
// 	x, y int
// }

type XYPointer struct {
	getX() int
	getY() int
}

// EuclidianPow2 - пусть расстояние между точками это Евклидова-метрика без ее корня
func EuclidianPow2(p1, p2 XYPoint) int {
	return (p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)
}

func main() {
	samples := []Point{XYPoint{0, 0}, XYPoint{10, 0}, XYPoint{10, 0}, XYPoint{10, 10}}
	fmt.Println(samples)
	c := Classificator{}
	fmt.Println(c)
	c.centroids = samples
	p := Point(XYPoint{1, 1})
	_ = p
	// Как EuclidianPow2 "подсунуть" теперь в getNearestCentroid
	fmt.Println(c.getNearestCentroid(p, EuclidianPow2))
}
