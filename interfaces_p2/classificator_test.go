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

func (p XYPoint) getCoordinates() interface{} {
	return [2]float64{p.x, p.y}
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
	classificator := Classificator{centroids}
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

func (p XYZPoint) getCoordinates() interface{} {
	return [3]float64{p.x, p.y, p.z}
}

// Minkovski - пусть расстояние между точками - это манхэттоновское
func Minkovski(p1, p2 Pointer) float64 {
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
	classificator := Classificator{centroids}
	fmt.Println("Classificator", classificator)
	point := Pointer(XYZPoint{x: 0, y: 0, z: 0})
	fmt.Println("Point", point)
	classifiedTo := classificator.getNearestCentroid(point, Minkovski)
	fmt.Println("ClassifiedTo", classifiedTo)
	// if classifiedTo != awaitedCentroid {
	if !Eq(classifiedTo, awaitedCentroid) {
		t.Errorf("Point %v must be classified to  %v, but was %v. ERROR.", point, awaitedCentroid, classifiedTo)
	} else {
		fmt.Printf("Point %v classified to  %v (%v). OK.\n", point, classifiedTo, awaitedCentroid)
	}
}
