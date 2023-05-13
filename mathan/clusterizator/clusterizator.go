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
//	Metrica - функция расстояния в пространстве.
//	Breakpoint - критерий останова, например, по числу итоговых кластеров:
//	   	func (c *Clusterizator) Breakpoint() bool {
//	   		if len(c.clusters) == 10 {
//	   			return true
//	   		}
//	   		return false
//	   	}
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
