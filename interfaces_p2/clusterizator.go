package main

// Clusterizator - кластеризатор.
type Clusterizator struct {
	clusters [][]Pointer
}

// Критерий останова
// func stopCriteria(c *Clusterizator) bool {
// 	if len(c.clusters) == 10 {
// 		return true
// 	}
// 	return false
// }

func Nearest(point Pointer, samples []Pointer, DistanceBetween func(a, b Pointer) float64) (float64, Pointer) {
	default_distance := float64(-1)
	distance := default_distance
	var unionTo Pointer
	for _, sample := range samples {
		pointToSampleDistance := DistanceBetween(p, sample)
		if distance == -1 {
			distance = pointToSampleDistance
			unionTo = sample
			continue
		}
		if pointToSampleDistance < distance {
			distance = pointToSampleDistance
			unionTo = sample
		}
	}
	// if distance == default_distance {
	// 	0, point
	// }
	return distance, unionTo
}


func (c *Clusterizator) clusterize(p Pointer, samples []Pointer, DistanceBetween func(a, b Pointer) float64) (float64, Pointer) {
	distance := float64(-1)
	var unionTo Pointer
	for i, candidateToCluster := range c.clusters[:len(c.clusters)-2]  {
		distance, clusterEthalon := Nearest(candidateToCluster, c.clusters[i+1:])
	}
	return distance, unionTo
}
