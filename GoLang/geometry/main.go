package main

import "math"

type Point struct {
	X, Y int
}
type Path []Point

// Distance returns the distance between two points.
func (p Point) Distance(q Point) int {
	return int(math.Hypot(float64(p.X-q.X), float64(p.Y-q.Y)))
}

func (p Path) Distance() int {
	sum := 0
	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}
