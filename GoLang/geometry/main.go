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

type IntList struct {
    Value int
    Tail *IntList
}

func (list *IntList) Sum() int {
    if list == nil {
        return 0
    }
    return list.Value + list.Tail.Sum()
}

func (list *IntList) Add(val int) {
    if list == nil {
        return
    }
    for list.Tail != nil {
        list = list.Tail
    }
    list.Tail = &IntList{val, nil}
}

func main(){
    var data IntList = IntList{1, &IntList{2, &IntList{3, nil}}}
    println(data.Sum())
    data.Add(4)
    println(data.Sum())
}
