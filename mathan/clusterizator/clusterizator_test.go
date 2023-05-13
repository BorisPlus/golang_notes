package clusterizator_test

import (
	"fmt"
	"testing"

	dlist "github.com/BorisPlus/golang_notes/dlist"
	ztr "github.com/BorisPlus/golang_notes/mathan/clusterizator"
	pnt "github.com/BorisPlus/golang_notes/mathan/pointer"
)

// Point2D - структура точки двумерного пространства.
type Point2D struct {
	x, y float64
}

// String() - реализация интерфейса Stringer.
func (p Point2D) String() string {
	return fmt.Sprintf("(%v;%v)", p.x, p.y)
}

// String() - реализация интерфейса Pointer.
func (p Point2D) Coordinates() interface{} {
	return [2]float64{p.x, p.y}
}

// Euclidian(p1, p2 pnt.Pointer) - пусть расстояние между точками это Евклидова-метрика без ее корня.
func Euclidian(p1, p2 pnt.Pointer) float64 {
	return ((p1.(Point2D).x-p2.(Point2D).x)*(p1.(Point2D).x-p2.(Point2D).x) +
		(p1.(Point2D).y-p2.(Point2D).y)*(p1.(Point2D).y-p2.(Point2D).y))
}

func avg(x1, x2 float64) float64 {
	return (x1 + x2) / 2
}

// Centroid(p1, p2 pnt.Pointer) - функция вычисления координаты "середины" между точками.
func Centroid(p1, p2 pnt.Pointer) pnt.Pointer {
	return Point2D{
		avg(p1.(Point2D).x, p2.(Point2D).x),
		avg(p1.(Point2D).y, p2.(Point2D).y),
	}
}

func TestClusterizatorPointValue(t *testing.T) {

	_ = dlist.DListItem{}
	tests := []struct {
		msg   string
		input []pnt.Pointer
	}{
		{
			msg: "4 points",
			input: []pnt.Pointer{
				Point2D{x: 0, y: 0},
				Point2D{x: 0, y: 1},
				Point2D{x: 10, y: 0},
				Point2D{x: 10, y: 2},
			},
		},
		{
			msg: "3 points",
			input: []pnt.Pointer{
				Point2D{x: 1, y: 0},
				Point2D{x: 0, y: 1},
				Point2D{x: 0.7, y: 0.5},
			},
		},
		{
			msg: "twice point",
			input: []pnt.Pointer{
				Point2D{x: 0, y: 0},
				Point2D{x: 0, y: 0},
			},
		},
		{
			msg: "1 point",
			input: []pnt.Pointer{
				Point2D{x: 1, y: 0},
			},
		},
		{
			msg: "No points",
			input: []pnt.Pointer{
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			clz := ztr.Clusterizator{
				Metrica:    Euclidian,
				Breakpoint: ztr.MustBeOne,
				Centroid:   Centroid,
			}
			clusters := clz.Init(tc.input).Clusterize()
			for _, cluster := range clusters {
				fmt.Println(cluster)
			}
		})
	}
}
