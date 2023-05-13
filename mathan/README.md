# Интерфейсы. Классификация и кластеризация данных

## Реализация

Пояснения представлены в рамках комментариев.

### Интерфейс точки

```go
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

```

### Классификация

Ниже представлен вариант реализации на Golang алгоритма соотнесения точки пространства с тем или иным классом точек, определяемым его центром, посредством интерфейса без привязки к метрике (функция расчёта расстояния между точками и размерность пространства задаются при задействовании указанного интерфейса и структуры классификатора).

Основной посыл, что классификатор, опираясь только на **интерфейс**, заранее не знает размерность пространства и функцию расчёта расстояния, но решает задачу классификации.

#### Код

<details>
<summary>см. "classificator.go"</summary>

```go
package classificator

import (
    pnt "github.com/BorisPlus/golang_notes/mathan/pointer"
)

// Решение задач классификации посредством интерфейса без привязки к метрике
// (без конкретной функции расчёта расстояния между точками и размерности пространства).

// Classificator - структура классификатора.
type Classificator struct {
    centroids []pnt.Pointer // центры заранее известных классов.
}

// SetCentroids(centroids []Pointer) - определение центров заранее известных классов.
func (c *Classificator) SetCentroids(centroids []pnt.Pointer) {
    c.centroids = centroids
}

// Classificator.getNearestCentroid - получает ближайший центр класса в зависимости от
// переданной функции расстояния - metrica (заранее не известной)
func (c *Classificator) Classify(point pnt.Pointer, metrica func(a, b pnt.Pointer) float64) pnt.Pointer {
    distance := float64(-1)
    var pointCentroid pnt.Pointer
    for _, centroid := range c.centroids {
        pointToSampleDistance := metrica(centroid, point)
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

</details>

#### Тестирование

<details>
<summary>см. "classificator_test.go"</summary>

```go
package classificator_test

import (
    "fmt"
    "math"
    "testing"

    pnt "github.com/BorisPlus/golang_notes/mathan/pointer"
    cls "github.com/BorisPlus/golang_notes/mathan/classificator"

)

// ДВУМЕРНОЕ ПРОСТРАНСТВО ===========================

// XYPoint - пусть точка Point двумерная
type XYPoint struct {
    x, y float64
}

func (p XYPoint) String() string {
    return fmt.Sprintf("(%v;%v)", p.x, p.y)
}

func (p XYPoint) Coordinates() interface{} {
    return [2]float64{p.x, p.y}
}

// EuclidianXYPow2 - пусть расстояние между точками это Евклидова-метрика без ее корня
func EuclidianXYPow2(p1, p2 pnt.Pointer) float64 {
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
    awaitedCentroid := pnt.Pointer(XYPoint{x: 0, y: 0})
    fmt.Println("AwaitedCentroid", awaitedCentroid)
    centroids := []pnt.Pointer{
        XYPoint{x: 10, y: 0},
        XYPoint{x: 0, y: 10},
        XYPoint{x: 10, y: 10}}
    //
    centroids = append(centroids, awaitedCentroid)
    fmt.Println("Centroids", centroids)
    classificator := cls.Classificator{}
    classificator.SetCentroids(centroids)
    fmt.Println("Classificator", classificator)
    // Берем точку с координатами
    point := pnt.Pointer(XYPoint{x: 1, y: 1})
    fmt.Println("Point", point)
    // Выясняем какой центройд ей ближе по Евклиду
    classifiedTo := classificator.Classify(point, EuclidianXYPow2)
    fmt.Println("ClassifiedTo", classifiedTo)
    if !pnt.Eq(classifiedTo, awaitedCentroid) {
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

func (p XYZPoint) Coordinates() interface{} {
    return [3]float64{p.x, p.y, p.z}
}

// Minkovski - пусть расстояние между точками - это манхэттоновское
func Minkovski(p1, p2 pnt.Pointer) float64 {
    dXdYdZ := math.Abs(p1.(XYZPoint).x-p2.(XYZPoint).x) +
        math.Abs(p1.(XYZPoint).y-p2.(XYZPoint).y) +
        math.Abs(p1.(XYZPoint).z-p2.(XYZPoint).z)
    return dXdYdZ
}

func TestXYZ(t *testing.T) {
    awaitedCentroid := pnt.Pointer(XYZPoint{x: 0, y: 0, z: 0})
    centroids := []pnt.Pointer{
        XYZPoint{x: 10, y: 0, z: 10},
        XYZPoint{x: 0, y: 10, z: 5},
        XYZPoint{x: 10, y: 10, z: 10}}
    centroids = append(centroids, awaitedCentroid)
    fmt.Println("Centroids", centroids)
    //
    classificator := cls.Classificator{}
    classificator.SetCentroids(centroids)
    fmt.Println("Classificator", classificator)
    point := pnt.Pointer(XYZPoint{x: 0, y: 0, z: 0})
    fmt.Println("Point", point)
    classifiedTo := classificator.Classify(point, Minkovski)
    fmt.Println("ClassifiedTo", classifiedTo)
    if !pnt.Eq(classifiedTo, awaitedCentroid) {
        t.Errorf("Point %v must be classified to  %v, but was %v. ERROR.", point, awaitedCentroid, classifiedTo)
    } else {
        fmt.Printf("Point %v classified to  %v (%v). OK.\n", point, classifiedTo, awaitedCentroid)
    }
}

```

</details>

#### Демонстрация

```shell
go test -v ./ > ./classificator.go.txt
```

Лог:

```text
=== RUN   TestXY
AwaitedCentroid (0;0)
Centroids [(10;0) (0;10) (10;10) (0;0)]
Classificator {[{10 0} {0 10} {10 10} {0 0}]}
Point (1;1)
ClassifiedTo (0;0)
Point (1;1) classified to  (0;0) ((0;0)). OK.
--- PASS: TestXY (0.00s)
=== RUN   TestXYZ
Centroids [(10;0;10) (0;10;5) (10;10;10) (0;0;0)]
Classificator {[{10 0 10} {0 10 5} {10 10 10} {0 0 0}]}
Point (0;0;0)
ClassifiedTo (0;0;0)
Point (0;0;0) classified to  (0;0;0) ((0;0;0)). OK.
--- PASS: TestXYZ (0.00s)
PASS
ok  	github.com/BorisPlus/golang_notes/mathan/classificator	(cached)

```

### Иерархическая кластеризация

Ниже представлен вариант реализации на Golang иерархического алгоритма группировки точек пространства в кластеры.

Основной посыл, что кластеризатор, опираясь только на **интерфейс**, заранее не знает:

* размерность пространства;
* функцию расчёта расстояния;
* функцию расчёта центра кластера, формируемого слиянием двух кластеров со своими центром каждый;
  
и при этом решает задачу классификации.

#### Код кластеризатора

Решение задачи иерархической кластеризации:

<details>
<summary>см. "clusterizator.go"</summary>

```go
package clusterizator

import (
    "fmt"
    "sort"
    "strings"

    dlist "github.com/BorisPlus/golang_notes/dlist"
    pnt "github.com/BorisPlus/golang_notes/mathan/pointer"
)

// Cluster - иерархический кластер.
type Cluster struct {
    centroid   pnt.Pointer
    branchA    *Cluster
    branchB    *Cluster
    distanceAB float64
}

// Init(centroid pnt.Pointer, branchA *Cluster, branchB *Cluster, distanceAB float64) -
// инициализация кластера всем сразу.
func (clt *Cluster) Init(centroid pnt.Pointer, branchA *Cluster, branchB *Cluster, distanceAB float64) {
    clt.centroid = centroid
    clt.branchA = branchA
    clt.branchB = branchB
    clt.distanceAB = distanceAB
}

// Centroid() - координата кластера, центроид.
func (clt *Cluster) Centroid() pnt.Pointer {
    return clt.centroid
}

// BranchA() - ветка дочернего подкластера условного-"A".
func (clt *Cluster) BranchA() *Cluster {
    return clt.branchA
}

// BranchB() - ветка дочернего подкластера условного-"B".
func (clt *Cluster) BranchB() *Cluster {
    return clt.branchB
}

// DistanceAB() - дистанция между дочерними подкластерами кластера.
func (clt *Cluster) DistanceAB() float64 {
    return clt.distanceAB
}

var clusterStringTemplate = `
╒=====================╕
|Cluster: %p|
├---------------------┤
|centroid:%s
|abDist:  %f
|branchA: %s
|branchB: %s
╘=====================`

func tabDecorate(clt *Cluster) string {
    if clt == nil {
        return "<nil>" // c.String() не работает в этом варианте, но линтер упорно его требует.
    }
    return "⤵\n|\t" + strings.Replace(clt.String(), string('\n'), "\n|\t", -1)
}

// String() - реализация интерфейса Stringer().
func (clt Cluster) String() string {
    return fmt.Sprintf(
        clusterStringTemplate,
        &clt,
        clt.Centroid(),
        clt.DistanceAB(),
        tabDecorate(clt.BranchA()),
        tabDecorate(clt.BranchB()),
    )
}

// SimilarClustersPair - очень сложная структура ИСКЛЮЧИТЕЛЬНО для внутреннего использования.
// Содержит пары кластеров и их объединение.
// Задействуется в хранении максимально близких пар кластеров.
// Необходима для реализации big-O(1) при:
//   - вставке нового кластера
//   - пересчера новых расстояний
//   - удаления обработанных кластеров из всех списков
type SimilarClustersPairItem struct {
    thisItem    *dlist.DListItem
    itemA       *dlist.DListItem
    itemB       *dlist.DListItem
    abPairUnion *Cluster
}

// ABPairUnion() - вид кластера, который будет получен в результате объединения пары кластеров.
func (scp *SimilarClustersPairItem) ABPairUnion() *Cluster {
    return scp.abPairUnion
}

// Distance() - растояние между центрами кластеров, образующих объединения.
func (scp *SimilarClustersPairItem) Distance() float64 {
    return scp.abPairUnion.DistanceAB()
}

// Clusterizator - кластеризатор, реализующий процесс Иерархической кластеризации.
//
//    Metrica - функция расстояния в пространстве.
//    Breakpoint - критерий останова, например, по числу итоговых кластеров:
//           func (c *Clusterizator) Breakpoint() bool {
//               if len(c.clusters) == 10 {
//                   return true
//               }
//               return false
//           }
//
// В пакете имеются типовые критерии останова.
type Clusterizator struct {
    Metrica          func(a, b pnt.Pointer) float64
    Breakpoint       func(c *Clusterizator) bool
    Centroid         func(x, y pnt.Pointer) pnt.Pointer
    clusters         *dlist.DList
    mappedSimilarest map[*dlist.DListItem]*SimilarClustersPairItem
    similarest       *dlist.DList
}

// Clusters() - кластеры.
func (clz *Clusterizator) Clusters() *dlist.DList {
    return clz.clusters
}

// InitSimilarest(similarest *dlist.DList) - инициализация списка максимально подобных кластеров.
func (clz *Clusterizator) InitSimilarest(similarest *dlist.DList) {
    clz.similarest = similarest
    clz.similarest.SetLessItems(func(x, y *dlist.DListItem) bool {
        return x.Value().(*SimilarClustersPairItem).Distance() < y.Value().(*SimilarClustersPairItem).Distance()
    })
}

// Similarest() - список пар максимально подобных кластеров.
func (clz *Clusterizator) Similarest() *dlist.DList {
    return clz.similarest
}

// Init(points []pnt.Pointer) - инициализация кластеризатора.
func (clz *Clusterizator) Init(points []pnt.Pointer) *Clusterizator {
    clz.clusters = &dlist.DList{}
    for _, point := range points {
        cluster := Cluster{}
        cluster.Init(point, nil, nil, 0)
        clz.clusters.PushToLeftEdge(&cluster)
    }
    clz.calculateSimilarest()
    return clz
}

// calculateSimilarest() - вычисление матрицы расстояний между кластерами.
func (clz *Clusterizator) calculateSimilarest() {
    clz.InitSimilarest(&dlist.DList{})

    clz.mappedSimilarest = make(map[*dlist.DListItem]*SimilarClustersPairItem, 0)
    // Вот тут злосчастный big-O(n^2), точнее big-O(n*(n-1)))
    for exampleItem := clz.clusters.LeftEdge(); exampleItem != clz.clusters.RightEdge(); exampleItem = exampleItem.RightNeighbour() {
        //
        exampleCluster := exampleItem.Value().(*Cluster)
        //
        for similarItem := exampleItem.RightNeighbour(); similarItem != nil; similarItem = similarItem.RightNeighbour() {
            similarCluster := similarItem.Value().(*Cluster)
            distance := clz.Metrica(exampleCluster.Centroid(), similarCluster.Centroid())
            esCluster := Cluster{}

            if (clz.mappedSimilarest[exampleItem] == nil) ||
                (clz.mappedSimilarest[exampleItem] != nil && distance < clz.mappedSimilarest[exampleItem].ABPairUnion().DistanceAB()) {
                esCluster.Init(
                    clz.Centroid(
                        exampleCluster.Centroid(),
                        similarCluster.Centroid(),
                    ),
                    exampleCluster,
                    similarCluster,
                    distance,
                )
                clz.mappedSimilarest[exampleItem] = &SimilarClustersPairItem{
                    thisItem:    nil,
                    itemA:       exampleItem,
                    itemB:       similarItem,
                    abPairUnion: &esCluster,
                }
            }
        }
    }
    for _, similarest := range clz.mappedSimilarest {
        similarestItem := clz.similarest.PushToLeftEdge(similarest)
        similarest.thisItem = similarestItem
    }
}

func (clz *Clusterizator) CandidatesOfMerging() *dlist.DListItem {
    // TODO: Valuer for Less()!!!
    sort.Sort(clz.similarest)
    // TODO: можно запараллелить для одинаковых расстояний разных пар кластеров
    // return clz.similarest.LeftEdge().Value().(*privateSimilarCluster)
    return clz.similarest.LeftEdge()
}

func (clz *Clusterizator) Iterate() (bool, error) {
    if clz.Breakpoint(clz) {
        return true, nil
    }
    // Выбранные кандидаты на слияние.
    similarClustersItem := clz.CandidatesOfMerging()
    similarClustersItemA := similarClustersItem.Value().(*SimilarClustersPairItem)
    // centroid := clz.CentroidCalculator(similarClustersItem.Value().(*Cluster).branchA.centroid, similarClustersItem.Value().(*Cluster).branchB.centroid)
    itemA := similarClustersItemA.itemA
    // eCluster := eClusterItem.Value().(*Cluster)
    itemB := similarClustersItemA.itemB
    // sCluster := sClusterItem.Value().(*Cluster)
    newCluster := similarClustersItemA.ABPairUnion()
    // newCluster() = ""
    // newCluster.Init(centroid, aCluster, bCluster, distance)
    // Вычисляем расстояния от кластеров (за исключением выбранных кандидатов) до нового кластера.
    // В рамках словарей похожих.
    clz.clusters.Remove(itemA)
    clz.clusters.Remove(itemB)
    newClusterItem := clz.clusters.PushToRightEdge(newCluster)
    // fmt.Println("Максимально подобная пара")
    // fmt.Println(newCluster)
    clz.similarest.Remove(similarClustersItem)

    for clusterItem := clz.clusters.LeftEdge(); clusterItem != clz.clusters.RightEdge(); clusterItem = clusterItem.RightNeighbour() {

        distance := clz.Metrica(clusterItem.Value().(*Cluster).centroid, newCluster.Centroid())
        esCluster := Cluster{}
        if (clz.mappedSimilarest[clusterItem] == nil) ||
            (clz.mappedSimilarest[clusterItem] != nil && distance < clz.mappedSimilarest[clusterItem].ABPairUnion().DistanceAB()) {

            if clz.mappedSimilarest[clusterItem] != nil {
                sm := clz.mappedSimilarest[clusterItem]
                clz.similarest.Remove(sm.thisItem)
            }

            esCluster.Init(
                clz.Centroid(
                    clusterItem.Value().(*Cluster).centroid,
                    newCluster.Centroid(),
                ),
                clusterItem.Value().(*Cluster),
                newCluster,
                distance,
            )
            clz.mappedSimilarest[clusterItem] = &SimilarClustersPairItem{
                thisItem:    nil,
                itemA:       clusterItem,
                itemB:       newClusterItem,
                abPairUnion: &esCluster,
            }
            similarestItem := clz.similarest.PushToLeftEdge(clz.mappedSimilarest[clusterItem])
            clz.mappedSimilarest[clusterItem].thisItem = similarestItem
        }
    }
    return clz.Breakpoint(clz), nil
}

// Clusterize() - метод запуска кластеризации.
func (clz *Clusterizator) Clusterize() []*Cluster {

    stoped := false
    for !stoped  {
        stoped, _ = clz.Iterate()
    }
    clusters := make([]*Cluster, clz.clusters.Len())

    for clusterItem := clz.clusters.LeftEdge(); clusterItem != nil; clusterItem = clusterItem.RightNeighbour() {
        clusters[len(clusters)-1] = clusterItem.Value().(*Cluster)
    }

    return clusters
}

```

</details>

Критерии останова (самые простые и без логического обоснования)

<details>
<summary>см. "breakpoints.go"</summary>

```go
package clusterizator

func CountLimit(clz *Clusterizator, count int) bool {
    if count < 1 {
        count = 1
    }
    return clz.Clusters().Len() <= count
}

func Limit10(clz *Clusterizator) bool {
    return CountLimit(clz, 10)
}

func MustBeOne(clz *Clusterizator) bool {
    return CountLimit(clz, 1)
}

```

</details>

#### Тестирование кластеризации

<details>
<summary>см. "clusterizator_test.go"</summary>

```go
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

```

</details>

#### Демонстрация кластеризации

```shell
go test -v ./ > ./clusterizator.go.txt
```

Лог:

```text
=== RUN   TestClusterizatorPointValue
=== RUN   TestClusterizatorPointValue/4_points

╒=====================╕
|Cluster: 0xc0000a88d0|
├---------------------┤
|centroid:(5;0.375)
|abDist:  25.062500
|branchA: ⤵
|    
|    ╒=====================╕
|    |Cluster: 0xc0000a8900|
|    ├---------------------┤
|    |centroid:(2.5;0.25)
|    |abDist:  25.250000
|    |branchA: ⤵
|    |    
|    |    ╒=====================╕
|    |    |Cluster: 0xc0000a8930|
|    |    ├---------------------┤
|    |    |centroid:(0;0.5)
|    |    |abDist:  1.000000
|    |    |branchA: ⤵
|    |    |    
|    |    |    ╒=====================╕
|    |    |    |Cluster: 0xc0000a8960|
|    |    |    ├---------------------┤
|    |    |    |centroid:(0;1)
|    |    |    |abDist:  0.000000
|    |    |    |branchA: <nil>
|    |    |    |branchB: <nil>
|    |    |    ╘=====================
|    |    |branchB: ⤵
|    |    |    
|    |    |    ╒=====================╕
|    |    |    |Cluster: 0xc0000a8990|
|    |    |    ├---------------------┤
|    |    |    |centroid:(0;0)
|    |    |    |abDist:  0.000000
|    |    |    |branchA: <nil>
|    |    |    |branchB: <nil>
|    |    |    ╘=====================
|    |    ╘=====================
|    |branchB: ⤵
|    |    
|    |    ╒=====================╕
|    |    |Cluster: 0xc0000a89c0|
|    |    ├---------------------┤
|    |    |centroid:(5;0)
|    |    |abDist:  100.000000
|    |    |branchA: ⤵
|    |    |    
|    |    |    ╒=====================╕
|    |    |    |Cluster: 0xc0000a89f0|
|    |    |    ├---------------------┤
|    |    |    |centroid:(10;0)
|    |    |    |abDist:  0.000000
|    |    |    |branchA: <nil>
|    |    |    |branchB: <nil>
|    |    |    ╘=====================
|    |    |branchB: ⤵
|    |    |    
|    |    |    ╒=====================╕
|    |    |    |Cluster: 0xc0000a8a20|
|    |    |    ├---------------------┤
|    |    |    |centroid:(0;0)
|    |    |    |abDist:  0.000000
|    |    |    |branchA: <nil>
|    |    |    |branchB: <nil>
|    |    |    ╘=====================
|    |    ╘=====================
|    ╘=====================
|branchB: ⤵
|    
|    ╒=====================╕
|    |Cluster: 0xc0000a8a50|
|    ├---------------------┤
|    |centroid:(7.5;0.5)
|    |abDist:  26.000000
|    |branchA: ⤵
|    |    
|    |    ╒=====================╕
|    |    |Cluster: 0xc0000a8a80|
|    |    ├---------------------┤
|    |    |centroid:(10;1)
|    |    |abDist:  4.000000
|    |    |branchA: ⤵
|    |    |    
|    |    |    ╒=====================╕
|    |    |    |Cluster: 0xc0000a8ab0|
|    |    |    ├---------------------┤
|    |    |    |centroid:(10;2)
|    |    |    |abDist:  0.000000
|    |    |    |branchA: <nil>
|    |    |    |branchB: <nil>
|    |    |    ╘=====================
|    |    |branchB: ⤵
|    |    |    
|    |    |    ╒=====================╕
|    |    |    |Cluster: 0xc0000a8ae0|
|    |    |    ├---------------------┤
|    |    |    |centroid:(10;0)
|    |    |    |abDist:  0.000000
|    |    |    |branchA: <nil>
|    |    |    |branchB: <nil>
|    |    |    ╘=====================
|    |    ╘=====================
|    |branchB: ⤵
|    |    
|    |    ╒=====================╕
|    |    |Cluster: 0xc0000a8b10|
|    |    ├---------------------┤
|    |    |centroid:(5;0)
|    |    |abDist:  100.000000
|    |    |branchA: ⤵
|    |    |    
|    |    |    ╒=====================╕
|    |    |    |Cluster: 0xc0000a8b40|
|    |    |    ├---------------------┤
|    |    |    |centroid:(10;0)
|    |    |    |abDist:  0.000000
|    |    |    |branchA: <nil>
|    |    |    |branchB: <nil>
|    |    |    ╘=====================
|    |    |branchB: ⤵
|    |    |    
|    |    |    ╒=====================╕
|    |    |    |Cluster: 0xc0000a8b70|
|    |    |    ├---------------------┤
|    |    |    |centroid:(0;0)
|    |    |    |abDist:  0.000000
|    |    |    |branchA: <nil>
|    |    |    |branchB: <nil>
|    |    |    ╘=====================
|    |    ╘=====================
|    ╘=====================
╘=====================
=== RUN   TestClusterizatorPointValue/3_points

╒=====================╕
|Cluster: 0xc0000a8d80|
├---------------------┤
|centroid:(0.425;0.625)
|abDist:  1.285000
|branchA: ⤵
|    
|    ╒=====================╕
|    |Cluster: 0xc0000a8db0|
|    ├---------------------┤
|    |centroid:(0;1)
|    |abDist:  0.000000
|    |branchA: <nil>
|    |branchB: <nil>
|    ╘=====================
|branchB: ⤵
|    
|    ╒=====================╕
|    |Cluster: 0xc0000a8de0|
|    ├---------------------┤
|    |centroid:(0.85;0.25)
|    |abDist:  0.340000
|    |branchA: ⤵
|    |    
|    |    ╒=====================╕
|    |    |Cluster: 0xc0000a8e10|
|    |    ├---------------------┤
|    |    |centroid:(0.7;0.5)
|    |    |abDist:  0.000000
|    |    |branchA: <nil>
|    |    |branchB: <nil>
|    |    ╘=====================
|    |branchB: ⤵
|    |    
|    |    ╒=====================╕
|    |    |Cluster: 0xc0000a8e40|
|    |    ├---------------------┤
|    |    |centroid:(1;0)
|    |    |abDist:  0.000000
|    |    |branchA: <nil>
|    |    |branchB: <nil>
|    |    ╘=====================
|    ╘=====================
╘=====================
=== RUN   TestClusterizatorPointValue/twice_point

╒=====================╕
|Cluster: 0xc0000a8f90|
├---------------------┤
|centroid:(0;0)
|abDist:  0.000000
|branchA: ⤵
|    
|    ╒=====================╕
|    |Cluster: 0xc0000a8fc0|
|    ├---------------------┤
|    |centroid:(0;0)
|    |abDist:  0.000000
|    |branchA: <nil>
|    |branchB: <nil>
|    ╘=====================
|branchB: ⤵
|    
|    ╒=====================╕
|    |Cluster: 0xc0000a8ff0|
|    ├---------------------┤
|    |centroid:(0;0)
|    |abDist:  0.000000
|    |branchA: <nil>
|    |branchB: <nil>
|    ╘=====================
╘=====================
=== RUN   TestClusterizatorPointValue/1_point

╒=====================╕
|Cluster: 0xc0000a90e0|
├---------------------┤
|centroid:(1;0)
|abDist:  0.000000
|branchA: <nil>
|branchB: <nil>
╘=====================
=== RUN   TestClusterizatorPointValue/No_points
--- PASS: TestClusterizatorPointValue (0.00s)
    --- PASS: TestClusterizatorPointValue/4_points (0.00s)
    --- PASS: TestClusterizatorPointValue/3_points (0.00s)
    --- PASS: TestClusterizatorPointValue/twice_point (0.00s)
    --- PASS: TestClusterizatorPointValue/1_point (0.00s)
    --- PASS: TestClusterizatorPointValue/No_points (0.00s)
PASS
ok      github.com/BorisPlus/golang_notes/mathan/clusterizator    (cached)

```

## Вывод

Интерфейс как составляющая часть инкапсуляции добавляет гибкости для будущих решений.

Приведенные выше измерения пространства могут быть перенесены на сравнения категорий, шутливо - "синего" и "тёплого", необходимо только придерживаться сигнатур анонимных функций и **интерфейса**.


