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
