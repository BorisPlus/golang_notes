package main

// Classificator - классификатор.
type Classificator struct {
	// centroids - это центры классов.
	centroids []Pointer
	// classes - это сами классы со всеми точками.
	classes map[Pointer][]Pointer
}

// Classificator.getNearestCentroid - получает ближайший центр класса в зависимости от
// переданной функции расстояния - DistanceBetween (заранее не известной)
func (c *Classificator) getNearestCentroid(point Pointer, DistanceBetween func(a, b Pointer) float64) Pointer {
	distance := float64(-1)
	var pointCentroid Pointer
	for _, centroid := range c.centroids {
		pointToSampleDistance := DistanceBetween(centroid, point)
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
