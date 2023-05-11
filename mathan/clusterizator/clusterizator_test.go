package clusterizator_test

import (
	"fmt"
	"testing"

	ztr "github.com/BorisPlus/golang_notes/mathan/clusterizator"
	pnt "github.com/BorisPlus/golang_notes/mathan/pointer"
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

// Euclidian(p1, p2 pnt.Pointer) - пусть расстояние между точками это Евклидова-метрика без ее корня
func Euclidian(p1, p2 pnt.Pointer) float64 {
	return ((p1.(Point2D).x-p2.(Point2D).x)*(p1.(Point2D).x-p2.(Point2D).x) +
		(p1.(Point2D).y-p2.(Point2D).y)*(p1.(Point2D).y-p2.(Point2D).y))
}

// Как будет рассчитываться координата центра кластеров
func avg(x1, x2 float64) float64 {
	return (x1 + x2) / 2
}

func CentroidCalculator(a, b *ztr.Cluster) (pnt.Pointer, error) {
	aCentroid, err := a.Centroid()
	if err != nil {
		return Point2D{}, err
	}
	bCentroid, err := b.Centroid()
	if err != nil {
		return Point2D{}, err
	}
	return Point2D{
		avg(aCentroid.(Point2D).x, bCentroid.(Point2D).x),
		avg(aCentroid.(Point2D).y, bCentroid.(Point2D).y),
	}, nil
}

//

// func TestClusterizator(t *testing.T) {
	
// 	//
// 	points := []pnt.Pointer{
// 		Point2D{x: 0, y: 0},
// 		Point2D{x: 0, y: 1},
// 		Point2D{x: 10, y: 0},
// 		Point2D{x: 10, y: 2}}

// 	fmt.Println("========================================================================")
// 	fmt.Println("Исходные точки")
// 	for _, obj := range points {
// 		fmt.Println(obj)
// 	}
// 	// 
// 	clz := ztr.Clusterizator{
// 		Metrica:            Euclidian,
// 		BreakpointChecker:  ztr.MustBeOne,
// 		CentroidCalculator: CentroidCalculator,
// 	}
// 	clz.Init(points)
// 	//
// 	fmt.Println("========================================================================")
// 	fmt.Println("Исходные Кластеры")
// 	for _, obj := range clz.Clusters() {
// 		fmt.Println(obj)
// 	}
// 	//
// 	fmt.Println("========================================================================")
// 	fmt.Println("Вектор максимальной схожести кластеров (минимумы строк матрица расстояний)")
// 	fmt.Println("MAP")
// 	for _, obj := range clz.SimilarestAsMap {
// 		fmt.Println(obj)
// 	}
// 	fmt.Println("========================================================================")
// 	fmt.Println("Вектор максимальной схожести кластеров (минимумы строк матрица расстояний)")
// 	fmt.Println("STRUCT")
// 	for _, obj := range clz.Similarest {
// 		fmt.Println(obj)
// 	}

// 	fmt.Println("========================================================================")
// 	fmt.Println("Кандидаты на слияние:")
// 	fmt.Println(clz.CandidatesOfMerging())
// 	//
// 	fmt.Println("========================================================================")
// 	fmt.Println("Шаг кластеризации:")
// 	fmt.Println("Кандидаты на слияние:")
// 	fmt.Println(clz.CandidatesOfMerging())
// 	stoped, err := clz.Iterate()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	if stoped == true {
// 		fmt.Println("Останов")
// 		return
// 	}
// 	fmt.Println("Новые кластеры")
// 	for _, obj := range clz.Clusters() {
// 		fmt.Println(obj)
// 	}
// 	fmt.Println(len(clz.Clusters()))
// 	fmt.Println(cap(clz.Clusters()))
// 	fmt.Println(stoped)
// 	fmt.Println("Шаг кластеризации:")
// 	fmt.Println("Кандидаты на слияние:")
// 	fmt.Println(clz.CandidatesOfMerging())
// 	fmt.Println(len(clz.Clusters()))
// 	fmt.Println(cap(clz.Clusters()))
// 	fmt.Println("Шаг кластеризации:")
// 	stoped, err = clz.Iterate()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	if stoped == true {
// 		fmt.Println("Останов")
// 		return
// 	}
// 	fmt.Println("Новые кластеры")
// 	fmt.Println("Кандидаты на слияние:")
// 	fmt.Println(clz.CandidatesOfMerging())
// 	stoped, err = clz.Iterate()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	if stoped == true {
// 		fmt.Println("Останов")
// 		for _, obj := range clz.Clusters() {
// 			fmt.Println(obj)
// 		}
// 		return
// 	}
// 	fmt.Println("Новые кластеры")
// 	for _, obj := range clz.Clusters() {
// 		fmt.Println(obj)
// 	}
// 	fmt.Println(clz.BreakpointChecker(&clz))
// 	fmt.Println(len(clz.Clusters()))
// 	fmt.Println(cap(clz.Clusters()))
// 	fmt.Println(stoped)
// }

func TestClusterizatorLoop(t *testing.T) {
	
	//
	points := []pnt.Pointer{
		Point2D{x: 0, y: 0},
		Point2D{x: 0, y: 1},
		Point2D{x: 10, y: 0},
		Point2D{x: 10, y: 2}}

	fmt.Println("========================================================================")
	fmt.Println("Исходные точки:")
	for _, obj := range points {
		fmt.Println(obj)
	}
	// 
	clz := ztr.Clusterizator{
		Metrica:            Euclidian,
		BreakpointChecker:  ztr.MustBeOne,
		CentroidCalculator: CentroidCalculator,
	}
	clz.Init(points)
	//
	fmt.Println("========================================================================")
	fmt.Println("Исходные Кластеры:")
	for _, obj := range clz.Clusters() {
		fmt.Println(obj)
	}
	// TODO: проверка останов до шага должен быть
	fmt.Println("========================================================================")
	fmt.Println("Кластеризация")
	for stoped := false; stoped != true; stoped, _ = clz.Iterate() {
		fmt.Println("  * шаг кластеризации...")
	}
	// 
	fmt.Println("========================================================================")
	fmt.Println("Итоговые кластеры:")
	for _, obj := range clz.Clusters() {
		fmt.Println(obj)
	}
}
