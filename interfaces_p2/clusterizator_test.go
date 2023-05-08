package main

import (
	"fmt"
	"testing"
)

// ДВУМЕРНОЕ ПРОСТРАНСТВО ===========================
// go test -v ./pointer.go ./clusterizator.go ./clusterizator_test.go
// XYPoint - пусть точка Point двумерная

type Point2D struct {
	x, y float64
}

func (p Point2D) String() string {
	return fmt.Sprintf("(%v;%v)", p.x, p.y)
}

func (p Point2D) Coordinates() interface{} {
	return [2]float64{p.x, p.y}
}

func (p Point2D) Add(additional XPointer) XPointer {
	return Point2D{
		p.x + additional.(Point2D).x,
		p.y + additional.(Point2D).y,
	}
}

func avg(x1, x2 float64) float64 {
	return (x1 + x2) / 2
}

func (p Point2D) Avg(with XPointer) XPointer {
	return Point2D{
		avg(p.x, with.(Point2D).x),
		avg(p.y, with.(Point2D).y),
	}
}

// EuclidianXYPow2 - пусть расстояние между точками это Евклидова-метрика без ее корня
func Euclidian(p1, p2 Pointer) float64 {
	return ((p1.(Point2D).x-p2.(Point2D).x)*(p1.(Point2D).x-p2.(Point2D).x) +
		(p1.(Point2D).y-p2.(Point2D).y)*(p1.(Point2D).y-p2.(Point2D).y))
}


func CountLimit(c *Clusterizator, count int) bool {
	if count < 1 {
		count = 1
	}
	return len(c.clusters) <= count
}

func Limit10(c *Clusterizator) bool {
	return CountLimit(c, 10)
}

func MustBeOne(c *Clusterizator) bool {
	return CountLimit(c, 1)
}



func TestClusterizator(t *testing.T) {
	points := []XPointer{
		Point2D{x: 0, y: 0},
		Point2D{x: 0, y: 1},
		Point2D{x: 10, y: 0},
		Point2D{x: 10, y: 2}}

	clz := Clusterizator{
		metrica:      Euclidian,
		stopCriteria: MustBeOne,
	}
	clz.Init(points)
	//
	fmt.Println("========================================================================")
	fmt.Println("Кластеры")
	for _, obj := range clz.clusters {
		fmt.Println(obj)
	}
	//
	fmt.Println("========================================================================")
	fmt.Println("Вектор максимальной схожести кластеров (минимумы строк матрица расстояний)")
	fmt.Println("similarestAsMap")
	for _, obj := range clz.similarestAsMap {
		fmt.Println(obj)
	}
	fmt.Println("========================================================================")
	fmt.Println("Вектор максимальной схожести кластеров (минимумы строк матрица расстояний)")
	fmt.Println("similarest")
	for _, obj := range clz.similarest {
		fmt.Println(obj)
	}
	
	fmt.Println("========================================================================")
	fmt.Println("Кандидаты на слияние:")
	fmt.Println(clz.mergeStep())
	//
	fmt.Println("========================================================================")
	fmt.Println("Шаг кластеризации:")
	fmt.Println("Кандидаты на слияние:")
	fmt.Println(clz.mergeStep())
	stoped := clz.step()
	fmt.Println("Новые кластеры")
	for _, obj := range clz.clusters {
		fmt.Println(obj)
	}
	fmt.Println(len(clz.clusters))
	fmt.Println(cap(clz.clusters))
	fmt.Println(stoped)
	fmt.Println("Шаг кластеризации:")
	fmt.Println("Кандидаты на слияние:")
	fmt.Println(clz.mergeStep())
	stoped = clz.step()
	fmt.Println("Новые кластеры")
	for _, obj := range clz.clusters {
		fmt.Println(obj)
	}
	fmt.Println(len(clz.clusters))
	fmt.Println(cap(clz.clusters))
	fmt.Println(stoped)
	fmt.Println("Шаг кластеризации:")
	fmt.Println("Кандидаты на слияние:")
	fmt.Println(clz.mergeStep())
	stoped = clz.step()
	fmt.Println("Новые кластеры")
	for _, obj := range clz.clusters {
		fmt.Println(obj)
	}
	fmt.Println(len(clz.clusters))
	fmt.Println(cap(clz.clusters))
	fmt.Println(stoped)
	// fmt.Println("Шаг кластеризации:")
	// fmt.Println(clz.mergeStep())
	// clz.step()
	// fmt.Println("Новые кластеры")
	// for _, obj := range clz.clusters {
	// 	fmt.Println(obj)
	// }
	// fmt.Println(len(clz.clusters))
	// fmt.Println(cap(clz.clusters))

}
