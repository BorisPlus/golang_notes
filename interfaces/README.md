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

Ниже представлен вариант реализации на Golang алгоритма соотнесения точки пространства с тем или иным классом точек, определяемым его центром, посредством интерфейса без привязки к метрике (без функции расчёта расстояния между точками и размерности пространства).

Основной посыл, что классификатор заранее не знает размерность пространства и функцию подсчета расстояния, имеется только **интерфейс**.

#### Код

```go
package classificator

import (
    pnt "github.com/BorisPlus/golang_notes/interfaces/pointer"
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

#### Тестирование

<details>
<summary>см. "classificator_test.go"</summary>

```go
package classificator_test

import (
    "fmt"
    "math"
    "testing"

    pnt "github.com/BorisPlus/golang_notes/interfaces/pointer"
    cls "github.com/BorisPlus/golang_notes/interfaces/classificator"

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
ok  	github.com/BorisPlus/golang_notes/interfaces/classificator	0.005s

```

### Кластеризация (БУДЕТ ДОРАБОТАНА ДО БОЛЕЕ УНИВЕРСАЛЬНОГО ВАРИАНТА ПОЗЖЕ)

#### Код

Решение задачи иерархической кластеризации:

<details>
<summary>см. "clusterizator.go"</summary>

```go
package clusterizator

import (
    "fmt"
    "sort"
    "strings"

    pnt "github.com/BorisPlus/golang_notes/interfaces/pointer"
)

// Cluster - иерархический кластер.
type Cluster struct {
    centroid           pnt.Pointer
    branchA            *Cluster
    branchB            *Cluster
    centroidCalculator *func(x, y *Cluster) (pnt.Pointer, error)
}

func (clt *Cluster) SetCentroid(centroid pnt.Pointer) {
    clt.centroid = centroid
}

func (clt *Cluster) Centroid() (pnt.Pointer, error) {
    if clt.centroid == nil {
        err := clt.centroidCalculatorProxy()
        if err != nil {
            return nil, nil
        }
    }
    return clt.centroid, nil
}

func (clt *Cluster) CentroidCalculator() *func(x, y *Cluster) (pnt.Pointer, error) {
    return clt.centroidCalculator
}

func (clt *Cluster) SetBranchA(cluster *Cluster) {
    clt.branchA = cluster
    clt.centroid = nil
}

func (clt *Cluster) BranchA() *Cluster {
    return clt.branchA
}

func (clt *Cluster) SetBranchB(cluster *Cluster) {
    clt.branchB = cluster
    clt.centroid = nil
}

func (clt *Cluster) BranchB() *Cluster {
    return clt.branchB
}

func (clt *Cluster) centroidCalculatorProxy() error {
    if clt.branchA == nil {
        return fmt.Errorf("clt.branchA is nil")
    }
    if clt.branchB == nil {
        return fmt.Errorf("clt.branchA is nil")
    }
    centroid, err := (*clt.CentroidCalculator())(clt.branchA, clt.branchB)
    clt.centroid = centroid
    return err
}

var clusterStringTemplate = `╒=====================╕
|Cluster: %p|
├---------------------┤
|centroid:%s
|branchA: %s
|branchB: %s
╘=====================`

func tabDecorate(clt *Cluster) string {
    if clt == nil {
        return "<nil>" // c.String() не работает в этом варианте, но линтер упорно его требует.
    }
    return "⤵\n|\t" + strings.Replace(clt.String(), string('\n'), "\n|\t", -1)
}

func (clt Cluster) String() string {
    centroid, _ := clt.Centroid()
    return fmt.Sprintf(
        clusterStringTemplate,
        &clt, centroid,
        tabDecorate(clt.BranchA()),
        tabDecorate(clt.BranchB()),
    )
}

// SimilarCluster - структура, описывающая максимально близкий\похожий по метрике целевой кластер к исходному.
type SimilarCluster struct {
    example  *Cluster
    similar  *Cluster
    distance float64
}

var similarClusterStringTemplate = `
=====================
Similar:  %p
---------------------
example:  %p %p
similar:  %p %p
distance: %f
=====================`

func (smlr SimilarCluster) String() string {
    return fmt.Sprintf(
        similarClusterStringTemplate,
        &smlr,
        smlr.example,
        *smlr.example,
        smlr.similar,
        *smlr.similar,
        smlr.distance,
    )
}

// SimilarClusterVector - используется для сортировки пар максимально близких кластеров,
// в целях нахождения пары для следующей итерации слияния
type SimilarClusterVector []*SimilarCluster

func (v SimilarClusterVector) Len() int           { return len(v) }
func (v SimilarClusterVector) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v SimilarClusterVector) Less(i, j int) bool { return v[i].distance < v[j].distance }

// Clusterizator - кластеризатор.
//
//    metrica - функция расстояния в пространстве.
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
    Metrica            func(a, b pnt.Pointer) float64
    BreakpointChecker  func(c *Clusterizator) bool
    CentroidCalculator func(x, y *Cluster) (pnt.Pointer, error)
    clusters           []*Cluster
    SimilarestAsMap    map[*Cluster]*SimilarCluster
    Similarest         []*SimilarCluster
}

// Clusters() - кластеры.
func (clz *Clusterizator) Clusters() []*Cluster {
    return clz.clusters
}

// Init(points []pnt.Pointer) - инициализация кластеризатора.
func (clz *Clusterizator) Init(points []pnt.Pointer) {
    clz.clusters = make([]*Cluster, len(points))
    for index, point := range points {
        clz.clusters[index] = &Cluster{centroidCalculator: &clz.CentroidCalculator}
        clz.clusters[index].SetCentroid(point)
    }
    clz.calculateSimilarest()
}

// CalculateDistinctMatrix - вычисление матрицы расстояний между кластерами.
func (clz *Clusterizator) calculateSimilarest() error {
    clz.SimilarestAsMap = make(map[*Cluster]*SimilarCluster)
    //
    for srcIndx, exampleCluster := range clz.clusters[:len(clz.clusters)-1] {
        // нижеследующее вычисление можно загорутинить
        for _, similarCluster := range clz.clusters[srcIndx+1:] {
            exampleOrigin, err := exampleCluster.Centroid()
            if err != nil {
                return err
            }
            similarOrigin, err := similarCluster.Centroid()
            if err != nil {
                return err
            }
            distance := clz.Metrica(exampleOrigin, similarOrigin)
            if (clz.SimilarestAsMap[exampleCluster] == nil) ||
                (clz.SimilarestAsMap[exampleCluster] != nil && distance < clz.SimilarestAsMap[exampleCluster].distance) {
                clz.SimilarestAsMap[exampleCluster] = &SimilarCluster{
                    example:  exampleCluster,
                    similar:  similarCluster,
                    distance: distance,
                }
            }
        }
    }
    //
    clz.Similarest = make([]*SimilarCluster, len(clz.clusters)-1)
    idx := 0
    for _, similarCluster := range clz.SimilarestAsMap {
        clz.Similarest[idx] = similarCluster
        idx++
    }
    return nil
}

func (clz *Clusterizator) CandidatesOfMerging() *SimilarCluster {
    sort.Sort(SimilarClusterVector(clz.Similarest))
    // TODO: можно запараллелить для одинаковых расстоний разных пар кластеров
    return clz.Similarest[0]
}

func (clz *Clusterizator) Iterate() (bool, error) {
    // Выбранные кандидаты на слияние.
    similarClusters := clz.CandidatesOfMerging()
    // Новый кластер из выбранных кандидатов.
    newCluster := Cluster{centroidCalculator: &clz.CentroidCalculator}
    newCluster.SetBranchA(similarClusters.example)
    newCluster.SetBranchB(similarClusters.similar)
    // Вычисляем расстояния от кластеров (за исключением выбранных кандидатов) до нового кластера.
    // В рамках словарей похожих.
    for _, cluster := range clz.clusters {
        if cluster == newCluster.branchA {
            clz.SimilarestAsMap[cluster] = nil
            continue
        }
        if cluster == newCluster.branchB {
            clz.SimilarestAsMap[cluster] = nil
            continue
        }

        clusterCentroid, err := cluster.Centroid()

        if err != nil {
            return false, err
        }

        newClusterCentroid, err := newCluster.Centroid()

        if err != nil {
            return false, err
        }

        distance := clz.Metrica(clusterCentroid, newClusterCentroid)

        if (clz.SimilarestAsMap[cluster] == nil) ||
            (clz.SimilarestAsMap[cluster] != nil && distance < clz.SimilarestAsMap[cluster].distance) {
            clz.SimilarestAsMap[cluster] = &SimilarCluster{
                example:  cluster,
                similar:  &newCluster,
                distance: distance,
            }
        }
    }

    // В рамках слайсов похожих.
    clz.Similarest = make([]*SimilarCluster, len(clz.clusters)-2)
    idx := 0
    for _, similarCluster := range clz.SimilarestAsMap {
        if similarCluster != nil {
            clz.Similarest[idx] = similarCluster
            idx++
        }
    }

    // ---------------------
    // TODO: this is big-O(n)!!!

    for index, clusterToPop := range clz.clusters {
        if clusterToPop == newCluster.BranchA() {
            clz.clusters = append(clz.clusters[:index], clz.clusters[index+1:]...)
            break
        }
    }
    for index, clusterToPop := range clz.clusters {
        if clusterToPop == newCluster.BranchB() {
            clz.clusters = append(clz.clusters[:index], clz.clusters[index+1:]...)
            break
        }
    }
    clz.clusters = append(clz.clusters, &newCluster)
    // ---------------------
    return clz.BreakpointChecker(clz), nil
}

// func (c *Clusterizator) clusterizeAndUpdate(p Pointer, DistanceBetween func(a, b Pointer) float64) (float64, Pointer) {
//     distance, cluster := c.clusterize(p, DistanceBetween)
//     cluster := Average(cluster, p)
//     return cluster
// }

```

</details>

#### Тестирование

<details>
<summary>см. "clusterizator_test.go"</summary>

```go
package clusterizator_test

import (
    "fmt"
    "testing"

    ztr "github.com/BorisPlus/golang_notes/interfaces/clusterizator"
    pnt "github.com/BorisPlus/golang_notes/interfaces/pointer"
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
    
//     //
//     points := []pnt.Pointer{
//         Point2D{x: 0, y: 0},
//         Point2D{x: 0, y: 1},
//         Point2D{x: 10, y: 0},
//         Point2D{x: 10, y: 2}}

//     fmt.Println("========================================================================")
//     fmt.Println("Исходные точки")
//     for _, obj := range points {
//         fmt.Println(obj)
//     }
//     // 
//     clz := ztr.Clusterizator{
//         Metrica:            Euclidian,
//         BreakpointChecker:  ztr.MustBeOne,
//         CentroidCalculator: CentroidCalculator,
//     }
//     clz.Init(points)
//     //
//     fmt.Println("========================================================================")
//     fmt.Println("Исходные Кластеры")
//     for _, obj := range clz.Clusters() {
//         fmt.Println(obj)
//     }
//     //
//     fmt.Println("========================================================================")
//     fmt.Println("Вектор максимальной схожести кластеров (минимумы строк матрица расстояний)")
//     fmt.Println("MAP")
//     for _, obj := range clz.SimilarestAsMap {
//         fmt.Println(obj)
//     }
//     fmt.Println("========================================================================")
//     fmt.Println("Вектор максимальной схожести кластеров (минимумы строк матрица расстояний)")
//     fmt.Println("STRUCT")
//     for _, obj := range clz.Similarest {
//         fmt.Println(obj)
//     }

//     fmt.Println("========================================================================")
//     fmt.Println("Кандидаты на слияние:")
//     fmt.Println(clz.CandidatesOfMerging())
//     //
//     fmt.Println("========================================================================")
//     fmt.Println("Шаг кластеризации:")
//     fmt.Println("Кандидаты на слияние:")
//     fmt.Println(clz.CandidatesOfMerging())
//     stoped, err := clz.Iterate()
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     if stoped == true {
//         fmt.Println("Останов")
//         return
//     }
//     fmt.Println("Новые кластеры")
//     for _, obj := range clz.Clusters() {
//         fmt.Println(obj)
//     }
//     fmt.Println(len(clz.Clusters()))
//     fmt.Println(cap(clz.Clusters()))
//     fmt.Println(stoped)
//     fmt.Println("Шаг кластеризации:")
//     fmt.Println("Кандидаты на слияние:")
//     fmt.Println(clz.CandidatesOfMerging())
//     fmt.Println(len(clz.Clusters()))
//     fmt.Println(cap(clz.Clusters()))
//     fmt.Println("Шаг кластеризации:")
//     stoped, err = clz.Iterate()
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     if stoped == true {
//         fmt.Println("Останов")
//         return
//     }
//     fmt.Println("Новые кластеры")
//     fmt.Println("Кандидаты на слияние:")
//     fmt.Println(clz.CandidatesOfMerging())
//     stoped, err = clz.Iterate()
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     if stoped == true {
//         fmt.Println("Останов")
//         for _, obj := range clz.Clusters() {
//             fmt.Println(obj)
//         }
//         return
//     }
//     fmt.Println("Новые кластеры")
//     for _, obj := range clz.Clusters() {
//         fmt.Println(obj)
//     }
//     fmt.Println(clz.BreakpointChecker(&clz))
//     fmt.Println(len(clz.Clusters()))
//     fmt.Println(cap(clz.Clusters()))
//     fmt.Println(stoped)
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

```

</details>

#### Демонстрация

```shell
go test -v ./ > ./clusterizator.go.txt
```

Лог:

```text
=== RUN   TestClusterizatorLoop
========================================================================
Исходные точки:
(0;0)
(0;1)
(10;0)
(10;2)
========================================================================
Исходные Кластеры:
╒=====================╕
|Cluster: 0xc0000724e0|
├---------------------┤
|centroid:(0;0)
|branchA: <nil>
|branchB: <nil>
╘=====================
╒=====================╕
|Cluster: 0xc000072510|
├---------------------┤
|centroid:(0;1)
|branchA: <nil>
|branchB: <nil>
╘=====================
╒=====================╕
|Cluster: 0xc000072540|
├---------------------┤
|centroid:(10;0)
|branchA: <nil>
|branchB: <nil>
╘=====================
╒=====================╕
|Cluster: 0xc000072570|
├---------------------┤
|centroid:(10;2)
|branchA: <nil>
|branchB: <nil>
╘=====================
========================================================================
Кластеризация
  * шаг кластеризации...
  * шаг кластеризации...
  * шаг кластеризации...
========================================================================
Итоговые кластеры:
╒=====================╕
|Cluster: 0xc000072630|
├---------------------┤
|centroid:(5;0.75)
|branchA: ⤵
|    ╒=====================╕
|    |Cluster: 0xc000072660|
|    ├---------------------┤
|    |centroid:(0;0.5)
|    |branchA: ⤵
|    |    ╒=====================╕
|    |    |Cluster: 0xc000072690|
|    |    ├---------------------┤
|    |    |centroid:(0;0)
|    |    |branchA: <nil>
|    |    |branchB: <nil>
|    |    ╘=====================
|    |branchB: ⤵
|    |    ╒=====================╕
|    |    |Cluster: 0xc0000726c0|
|    |    ├---------------------┤
|    |    |centroid:(0;1)
|    |    |branchA: <nil>
|    |    |branchB: <nil>
|    |    ╘=====================
|    ╘=====================
|branchB: ⤵
|    ╒=====================╕
|    |Cluster: 0xc0000726f0|
|    ├---------------------┤
|    |centroid:(10;1)
|    |branchA: ⤵
|    |    ╒=====================╕
|    |    |Cluster: 0xc000072720|
|    |    ├---------------------┤
|    |    |centroid:(10;0)
|    |    |branchA: <nil>
|    |    |branchB: <nil>
|    |    ╘=====================
|    |branchB: ⤵
|    |    ╒=====================╕
|    |    |Cluster: 0xc000072750|
|    |    ├---------------------┤
|    |    |centroid:(10;2)
|    |    |branchA: <nil>
|    |    |branchB: <nil>
|    |    ╘=====================
|    ╘=====================
╘=====================
--- PASS: TestClusterizatorLoop (0.00s)
PASS
ok      github.com/BorisPlus/golang_notes/interfaces/clusterizator    0.004s

```

## Вывод

Интерфейс как составляющая часть инкапсуляции добавляет гибкости для будущих решений.

Приведенные выше измерения пространства могут быть перенесены на сравнения категорий, шутливо - "синего" и "тёплого", необходимо только придерживаться сигнатур анонимных функций и **интерфейса**.


