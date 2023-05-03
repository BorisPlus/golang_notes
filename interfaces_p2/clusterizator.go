package main

import (
	"fmt"
)

// Clusterizator - кластеризатор.
type Clusterizator struct {
	// clasters - это центры кластеров.
	clusters []Pointer
	// clastersMap - это сами кластеры со всеми точками.
	clustersMap map[Pointer][]Pointer
}

// Критерий останова
func (c *Clusterizator) stopCriteria() bool {
	if len(c.clusters) == 10 {
		return true
	}
	return false
}

// 
func (c *Clusterizator) clusterize(DistanceFunc func(a, b Pointer) float64) (float64, Pointer) {
	distance := float64(-1)
	var unionTo Pointer
	for i, candidateToCluster := range c.clusters[:len(c.clusters)-2]  {
		for _, clusterEthalon := range c.clusters[i+1:] {
			pointToSampleDistance := DistanceFunc(candidateToCluster, clusterEthalon)
			if distance == -1 {
				distance = pointToSampleDistance
				unionTo = clusterEthalon
				continue
			}
			if pointToSampleDistance < distance {
				distance = pointToSampleDistance
				unionTo = clusterEthalon
			}
		}
	}
	return distance, unionTo
}

// Добавляет точку в тот или иной класс классификатора
// func (c *Classificator) classify(point Pointer, DistanceFunc func(a, b Pointer) float64) {
// 	pointCentroid := c.getNearestCentroid(point, DistanceFunc)
// 	c.classes[pointCentroid] = append(c.classes[pointCentroid], point)
// }
