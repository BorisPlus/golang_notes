package main

import (
	"fmt"
)

type AbstractPoint interface {
	Distance(p AbstractPoint) int
}

type Point interface {
	AbstractPoint
	Distance(p AbstractPoint) int
}

type Classificator struct {
	samples []AbstractPoint
	classes map[AbstractPoint][]AbstractPoint
}

func (c *Classificator) samplerize(p AbstractPoint) AbstractPoint {
	distance := -1
	var pointSample AbstractPoint
	for _, sample := range c.samples {
		pointToSampleDistance := p.Distance(sample)
		if distance == -1 {
			distance = pointToSampleDistance
			pointSample = sample
			continue
		}
		if pointToSampleDistance < distance {
			distance = pointToSampleDistance
			pointSample = sample
		}
	}
	return pointSample
}

func (c *Classificator) classify(p Point) {
	pointSample := c.samplerize(p)
	c.classes[pointSample] = append(c.classes[pointSample], p)
}

type XYPoint struct {
	x, y int
}

func (p XYPoint) Distance(other XYPoint) int {
	return (p.x-other.x)*(p.x-other.x) + (p.y-other.y)*(p.y-other.y)
}

func main() {
	samples := []AbstractPoint{XYPoint{0, 0}, XYPoint{10, 0}, XYPoint{10, 0}, XYPoint{10, 10}}
	fmt.Println(samples)
	c := Classificator{}
	fmt.Println(c)
	c.samples = samples
	p := AbstractPoint(XYPoint{1, 1})
	_ = p
	// fmt.Println(c.samplerize(p))
}
