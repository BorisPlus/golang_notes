package main

import (
	"fmt"
	"sort"
	"strings"
)

type XPointer interface {
	Pointer
	Add(p XPointer) XPointer
	Avg(p XPointer) XPointer
}

// Cluster - иерархический кластер.
type Cluster struct {
	origin        XPointer // может лучше сделать указатель?
	clusterSource *Cluster
	clusterTarget *Cluster
}

var clusterStringTemplate = `
=====================
Cluster: %p
---------------------
origin:  %s
source:  %s
target:  %s
=====================`

func tabDecorate(c *Cluster) string {
	if c == nil {
		return fmt.Sprintf("%s", c) // c.String() не работает в этом варианте, но линтер упорно его требует.
	}
	return "⤵" + strings.Replace(c.String(), string('\n'), "\n\t", -1)
}

func (c Cluster) String() string {
	return fmt.Sprintf(clusterStringTemplate, &c, c.origin, tabDecorate(c.clusterSource), tabDecorate(c.clusterTarget))
}

// SimilarCluster - структура, описывающая максимально близкий\похожий по метрике целевой кластер к исходному.
type SimilarCluster struct {
	source   *Cluster
	target   *Cluster
	distance float64
}

var similarClusterStringTemplate = `
=====================
Similar: %p
---------------------
source:  %p %p
target:  %p %p
distance:%f
=====================`

func (s SimilarCluster) String() string {
	return fmt.Sprintf(similarClusterStringTemplate, &s, s.source, *s.source, s.target, *s.target, s.distance)
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
//	stopCriteria - критерий останова, например, по числу итоговых кластеров:
//	   	func (c *Clusterizator) StopCriteria() bool {
//	   		if len(c.clusters) == 10 {
//	   			return true
//	   		}
//	   		return false
//	   	}
type Clusterizator struct {
	metrica      func(a, b Pointer) float64
	stopCriteria func(c *Clusterizator) bool
	clusters     []*Cluster
	// similarestAsMap     map[*Cluster]map[*Cluster]float64
	similarestAsMap map[*Cluster]*SimilarCluster
	similarest      []*SimilarCluster
}

// Init - инициализация кластеризатора.
func (clz *Clusterizator) Init(points []XPointer) {
	clz.clusters = make([]*Cluster, len(points))
	for index, point := range points {
		clz.clusters[index] = &Cluster{point, nil, nil}
	}
	clz.calculateSimilarest()
}

// CalculateDistinctMatrix - вычисление матрицы расстояний кластеров.
func (clz *Clusterizator) calculateSimilarest() {
	clz.similarestAsMap = make(map[*Cluster]*SimilarCluster)
	//
	for srcIndx, sourceCluster := range clz.clusters[:len(clz.clusters)-1] {
		// нижеследующее вычисление можно загорутинить
		for _, targetCluster := range clz.clusters[srcIndx+1:] {
			distance := clz.metrica(sourceCluster.origin, targetCluster.origin)
			if (clz.similarestAsMap[sourceCluster] == nil) ||
				(clz.similarestAsMap[sourceCluster] != nil && distance < clz.similarestAsMap[sourceCluster].distance) {
				clz.similarestAsMap[sourceCluster] = &SimilarCluster{
					source:   sourceCluster,
					target:   targetCluster,
					distance: distance,
				}
			}
		}
	}
	//
	clz.similarest = make([]*SimilarCluster, len(clz.clusters)-1)
	idx := 0
	for _, similarCluster := range clz.similarestAsMap {
		clz.similarest[idx] = similarCluster
		idx++
	}
}

func (clz *Clusterizator) mergeStep() *SimilarCluster {
	sort.Sort(SimilarClusterVector(clz.similarest))
	// TODO: можно запараллелить для одинаковых расстоний разных пар кластеров
	return clz.similarest[0]
}

func (clz *Clusterizator) step() bool {
	// Выбранные кандидаты на слияние.
	similarClusters := clz.mergeStep()
	// Новый кластер из выбранных кандидатов.
	newCluster := Cluster{
		origin:        similarClusters.source.origin.Avg(similarClusters.target.origin),
		clusterSource: similarClusters.source, 
		clusterTarget: similarClusters.target, 
	}
	// Вычисляем расстояния от кластеров (за исключением выбранных кандидатов) до нового кластера.
	fmt.Println("Recalculate similarest")

	// В рамках словарей похожих.
	for _, cluster := range clz.clusters {
		if cluster == newCluster.clusterSource {
			clz.similarestAsMap[cluster] = nil
			continue
		}
		if cluster == newCluster.clusterTarget {
			clz.similarestAsMap[cluster] = nil
			continue
		}

		fmt.Println(cluster.origin, "<--->", newCluster.origin)
		fmt.Println("Для")
		fmt.Println(cluster)
		fmt.Println("Был")
		fmt.Println(clz.similarestAsMap[cluster])
		// fmt.Println(clz.similarestAsMap[cluster].source.origin)

		distance := clz.metrica(
			cluster.origin,
			newCluster.origin,
		)
		if (clz.similarestAsMap[cluster] == nil) ||
			(clz.similarestAsMap[cluster] != nil && distance < clz.similarestAsMap[cluster].distance) {
			clz.similarestAsMap[cluster] = &SimilarCluster{
				source:   cluster,
				target:   &newCluster,
				distance: distance,
			}
			fmt.Println("Стал")
			fmt.Println(clz.similarestAsMap[cluster])
		}
	}
	// clz.similarestAsMap[newCluster.clusterSource] = nil
	// clz.similarestAsMap[newCluster.clusterTarget] = nil

	// В рамках слайсов похожих.
	clz.similarest = make([]*SimilarCluster, len(clz.clusters)-2)
	idx := 0
	for _, similarCluster := range clz.similarestAsMap {
		if similarCluster != nil {
			clz.similarest[idx] = similarCluster
			idx++
		}
	}

	// ---------------------
	// TODO: this is big-O(n)!!!

	for index, clusterToPop := range clz.clusters {
		if clusterToPop == newCluster.clusterSource {
			clz.clusters = append(clz.clusters[:index], clz.clusters[index+1:]...)
			break
		}
	}
	for index, clusterToPop := range clz.clusters {
		if clusterToPop == newCluster.clusterTarget {
			clz.clusters = append(clz.clusters[:index], clz.clusters[index+1:]...)
			break
		}
	}
	clz.clusters = append(clz.clusters, &newCluster)

	// ---------------------


	// ---------------------
	// TODO: the same operations

	//
	// ---------------------
	return clz.stopCriteria(clz)
}

// func (c *Clusterizator) clusterizeAndUpdate(p Pointer, DistanceBetween func(a, b Pointer) float64) (float64, Pointer) {
// 	distance, cluster := c.clusterize(p, DistanceBetween)
// 	cluster := Average(cluster, p)
// 	return cluster
// }
