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
				Point2D{x: 0.5, y: 0.5},
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

func TeestClusterizatorLoop(t *testing.T) {

	_ = dlist.DListItem{}

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
	// //
	clz := ztr.Clusterizator{
		Metrica:    Euclidian,
		Breakpoint: ztr.MustBeOne,
		Centroid:   Centroid,
	}
	clz.Init(points)
	// //
	// fmt.Println("========================================================================")
	// fmt.Println("Исходные Кластеры:")
	// for cl := clz.Clusters().LeftEdge(); cl.RightNeighbour() != nil; cl = cl.RightNeighbour() {
	// 	fmt.Println(cl)
	// }

	// fmt.Println("========================================================================")
	fmt.Println("Максимально подобные Кластеры:")

	fmt.Printf("всего их %d\n", clz.Similarest().Len())

	if clz.Similarest().Len() != 3 {
		fmt.Errorf("clz.Similarest().Len() != 3, but %d", clz.Similarest().Len())
	}
	for cl := clz.Similarest().LeftEdge(); cl != nil; cl = cl.RightNeighbour() {
		fmt.Println("-----------------------------------------")
		fmt.Println(cl.Value())
		fmt.Println(cl.Value().(*ztr.SimilarClustersPairItem).ABPairUnion())
	}

	// clz.Similarest().SetValuer(func(v interface{}) float64 {
	// 	return v.(*ztr.SimilarItimedClusters).ESDistance()
	// })

	fmt.Println("-----------------------------------------")
	fmt.Println("Максимально подобная пара:")
	similarClustersItem := clz.CandidatesOfMerging()
	fmt.Println(similarClustersItem)

	// TODO: проверка останов до шага должен быть
	// fmt.Println("========================================================================")
	// fmt.Println("Кластеризация изнутри")
	// stoped, _ := clz.Iterate()
	// fmt.Println(stoped)

	// fmt.Println(clz.Clusters().Len())

	// fmt.Println("========================================================================")
	// fmt.Println("Новые Кластеры:")
	// for cl := clz.Clusters().LeftEdge(); cl != nil; cl = cl.RightNeighbour() {
	// 	fmt.Println(cl)
	// }
	// fmt.Println("Новые Максимально подобные Кластеры:")
	// for cl := clz.Similarest().LeftEdge(); cl != nil; cl = cl.RightNeighbour() {
	// 	fmt.Println(cl)
	// }

	fmt.Println("========================================================================")
	fmt.Println("Кластеризация (инкапсулировано)")
	for stoped, _ := clz.Iterate(); stoped != true; stoped, _ = clz.Iterate() {
		fmt.Println("-----------------------------------------")
		fmt.Println("  * шаг кластеризации...")
		fmt.Println("Новые Кластеры:")
		for cl := clz.Clusters().LeftEdge(); cl != nil; cl = cl.RightNeighbour() {
			fmt.Println(cl)
		}
		fmt.Println("Новые Максимально подобные Кластеры:")
		for cl := clz.Similarest().LeftEdge(); cl != nil; cl = cl.RightNeighbour() {
			fmt.Println(cl)
		}
	}

	// //
	// similarClustersItemA := similarClustersItem.Value().(struct {
	// 	exampleItem *dlist.DListItem
	// 	similarItem *dlist.DListItem
	// 	esCluster   *ztr.Cluster
	// })

	// eClusterItem := similarClustersItemA.exampleItem
	// fmt.Println(eClusterItem)
	// sClusterItem := similarClustersItemA.similarItem
	// fmt.Println(sClusterItem)
	// newCluster := similarClustersItemA.esCluster
	// fmt.Println(newCluster)

	// fmt.Println("========================================================================")
	// fmt.Println("Итоговые кластеры:")
	// for cl := clz.Clusters().LeftEdge();cl.RightNeighbour() !=nil;cl = cl.RightNeighbour() {
	// 	fmt.Println(cl)
	// }
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

// func TestClusterizatorLoop(t *testing.T) {

// 	_ = dlist.DListItem{}

// 	//
// 	points := []pnt.Pointer{
// 		Point2D{x: 0, y: 0},
// 		Point2D{x: 0, y: 1},
// 		Point2D{x: 10, y: 0},
// 		Point2D{x: 10, y: 2}}

// 	// fmt.Println("========================================================================")
// 	// fmt.Println("Исходные точки:")
// 	// for _, obj := range points {
// 	// 	fmt.Println(obj)
// 	// }
// 	// //
// 	clz := ztr.Clusterizator{
// 		Metrica:            Euclidian,
// 		BreakpointChecker:  ztr.MustBeOne,
// 		CentroidCalculator: CentroidCalculator,
// 	}
// 	clz.Init(points)
// 	// //
// 	// fmt.Println("========================================================================")
// 	// fmt.Println("Исходные Кластеры:")
// 	// for cl := clz.Clusters().LeftEdge(); cl.RightNeighbour() != nil; cl = cl.RightNeighbour() {
// 	// 	fmt.Println(cl)
// 	// }

// 	// fmt.Println("========================================================================")
// 	fmt.Println("Максимально подобные Кластеры:")

// 	fmt.Printf("всего их %d\n", clz.Similarest().Len())

// 	if clz.Similarest().Len() != 3 {
// 		fmt.Errorf("clz.Similarest().Len() != 3, but %d", clz.Similarest().Len())
// 	}
// 	for cl := clz.Similarest().LeftEdge(); cl != nil; cl = cl.RightNeighbour() {
// 		fmt.Println("-----------------------------------------")
// 		fmt.Println(cl.Value())
// 		fmt.Println(cl.Value().(*ztr.SimilarItimedClusters).ESCluster())
// 	}

// 	clz.Similarest().SetValuer(func(v interface{}) float64 {
// 		return v.(*ztr.SimilarItimedClusters).ESDistance()
// 	})

// 	fmt.Println("-----------------------------------------")
// 	fmt.Println("Максимально подобная пара:")
// 	similarClustersItem := clz.CandidatesOfMerging()
// 	fmt.Println(similarClustersItem)

// 	// TODO: проверка останов до шага должен быть
// 	// fmt.Println("========================================================================")
// 	// fmt.Println("Кластеризация изнутри")
// 	// stoped, _ := clz.Iterate()
// 	// fmt.Println(stoped)

// 	// fmt.Println(clz.Clusters().Len())

// 	// fmt.Println("========================================================================")
// 	// fmt.Println("Новые Кластеры:")
// 	// for cl := clz.Clusters().LeftEdge(); cl != nil; cl = cl.RightNeighbour() {
// 	// 	fmt.Println(cl)
// 	// }
// 	// fmt.Println("Новые Максимально подобные Кластеры:")
// 	// for cl := clz.Similarest().LeftEdge(); cl != nil; cl = cl.RightNeighbour() {
// 	// 	fmt.Println(cl)
// 	// }

// 	fmt.Println("========================================================================")
// 	fmt.Println("Кластеризация (инкапсулировано)")
// 	for stoped, _ := clz.Iterate(); stoped != true; stoped, _ = clz.Iterate() {
// 		fmt.Println("-----------------------------------------")
// 		fmt.Println("  * шаг кластеризации...")
// 		fmt.Println("Новые Кластеры:")
// 		for cl := clz.Clusters().LeftEdge(); cl != nil; cl = cl.RightNeighbour() {
// 			fmt.Println(cl)
// 		}
// 		fmt.Println("Новые Максимально подобные Кластеры:")
// 		for cl := clz.Similarest().LeftEdge(); cl != nil; cl = cl.RightNeighbour() {
// 			fmt.Println(cl)
// 		}
// 	}

// 	// //
// 	// similarClustersItemA := similarClustersItem.Value().(struct {
// 	// 	exampleItem *dlist.DListItem
// 	// 	similarItem *dlist.DListItem
// 	// 	esCluster   *ztr.Cluster
// 	// })

// 	// eClusterItem := similarClustersItemA.exampleItem
// 	// fmt.Println(eClusterItem)
// 	// sClusterItem := similarClustersItemA.similarItem
// 	// fmt.Println(sClusterItem)
// 	// newCluster := similarClustersItemA.esCluster
// 	// fmt.Println(newCluster)

// 	// fmt.Println("========================================================================")
// 	// fmt.Println("Итоговые кластеры:")
// 	// for cl := clz.Clusters().LeftEdge();cl.RightNeighbour() !=nil;cl = cl.RightNeighbour() {
// 	// 	fmt.Println(cl)
// 	// }
// }
