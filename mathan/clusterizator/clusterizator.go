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
//	metrica - функция расстояния в пространстве.
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
// 	distance, cluster := c.clusterize(p, DistanceBetween)
// 	cluster := Average(cluster, p)
// 	return cluster
// }
