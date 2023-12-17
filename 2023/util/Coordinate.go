package util

import (
	"strconv"
)

type Coordinate struct {
	X int
	Y int
}

type Direction int

// up = 0, down = 1 etc
const (
	Up Direction = iota
	Down
	Left
	Right
)

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

func (d Direction) Reverse() Direction {
	switch d {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	}
	panic("Please choose a valid direction")
}

func (d Direction) Turn(turn Direction) Direction {
	if turn != Left && turn != Right {
		panic("Only left or right")
	}

	switch d {
	case Up:
		return turn
	case Down:
		switch turn {
		case Left:
			return Right
		case Right:
			return Left
		}
	case Left:
		switch turn {
		case Left:
			return Down
		case Right:
			return Up
		}
	case Right:
		switch turn {
		case Left:
			return Up
		case Right:
			return Down
		}
	}

	panic("Please choose a valid direction")
}

func (c Coordinate) Move(direction Direction, moves int) Coordinate {
	switch direction {
	case Up:
		c.Y -= moves
	case Down:
		c.Y += moves
	case Left:
		c.X -= moves
	case Right:
		c.X += moves
	}
	return c
}
