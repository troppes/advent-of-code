package util

import (
	"strconv"
)

type Coordinate struct {
	X int
	Y int
}

func (a *Coordinate) Compare(b *Coordinate) bool {
	if a == b {
		return true
	}
	if a.X != b.X {
		return false
	}
	if a.Y != b.Y {
		return false
	}
	return true
}

func (c *Coordinate) String() string {
	return strconv.Itoa(c.X) + "," + strconv.Itoa(c.Y)
}

func (p *Coordinate) ManhattanDistance(q *Coordinate) int {
	return Abs(p.X-q.X) + Abs(p.Y-q.Y)
}
