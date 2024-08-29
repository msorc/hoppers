// Package hoppers provides functionality to find shortest path for Hoppers game
package hoppers

import (
	"errors"
	"fmt"
	"slices"
)

const (
	maxSpeed = 3
)

type Point struct {
	X int
	Y int
}

// Moves the point by one step applying the velocity.
// Movement is not performed with error if velicity is zero, point is inside obstacle,
//   moves outside of the grid, moves into obstacle, already visited
func (p Point) move(v velocity, grid grid) (Point, error) {
	if grid[p.X][p.Y].o {
		return p, errors.New("Inside obstacle")
	}

	newX := p.X + v.x
	newY := p.Y + v.y
	maxX := len(grid) - 1
	maxY := len(grid[0]) - 1

	if newX < 0 || newX > maxX || newY < 0 || newY > maxY {
		return p, errors.New("Out of border")
	}

	// Obstacle
	if grid[newX][newY].o {
		return p, errors.New("Obstacle")
	}
	// Visited before
	if slices.Contains(grid[newX][newY].vs, v) {
		return p, errors.New("Already visited")
	}

	return Point{newX, newY}, nil
}

type velocity struct {
	x int
	y int
}

// Returns a new velocity with acceleration on x and y applied.
// If absolute value exceeds maxSpeed, error is returnd.
func (v velocity) accelerate(x, y int) (velocity, error) {
	newX := v.x + x
	newY := v.y + y
	if abs(newX) > maxSpeed || abs(newY) > maxSpeed {
		return v, errors.New("Speed limit exceeded")
	}

	return velocity{newX, newY}, nil
}

// Check both axes acceleration is 0
func (v velocity) noMovement() bool {
	if v.x == 0 && v.y == 0 { return true }
	return false
}

type hop struct {
	p Point
	v velocity
}

type grid [][]cell

type cell struct {
	o bool // is obstacle
	vs []velocity // with which velocities the cell was visited
}

// Mark a cell as visited 
func (g grid) visit(h hop) {
	g[h.p.X][h.p.Y].vs = append(g[h.p.X][h.p.Y].vs, h.v)
}

func HopsCount(w, h int, start, finish Point, obstacles [][2]Point) (int, error) {
	grid := prepareGrid(w, h, obstacles)
	hops := []hop{hop{start, velocity{0, 0}}}

	hopCount, err := makeStep(hops, finish, grid, 0)
	if err != nil {
		return -1, errors.New("No solution")
	}

	return hopCount, nil
}

func makeStep(hops []hop, finish Point, grid grid, stepsTaken int) (int, error) {
	var newHops []hop

	for _, h := range hops {
		if h.p == finish {
			return stepsTaken, nil
		}
		grid.visit(h)

		pointHops := hopsFrom(h, grid)

		newHops = append(newHops, pointHops...)
	}

	if len(newHops) == 0 {
		return stepsTaken, errors.New("Can't move")
	}

	return makeStep(newHops, finish, grid, stepsTaken + 1)
}

func PresentResult(result int, err error) string {
	var output string

	if err != nil {
		output = err.Error()
	} else {
		output = fmt.Sprintf("Optimal solution takes %d hop(s)", result)
	}

	return output
}

// Allocates the grid structure and fills obstacles
func prepareGrid(w, h int, obstacles [][2]Point) grid {
	grid := make(grid, w)
	for i := range grid {
		grid[i] = make([]cell, h)
	}

	for _, o := range obstacles {
		for x := o[0].X; x <= o[1].X; x++ {
			for y := o[0].Y; y <= o[1].Y; y++ {
				grid[x][y].o = true
			}
		}
	}

	return grid
}

// Returns the list of hops from the current hop
func hopsFrom(h hop, grid grid) []hop {
	var speedChange = [3]int{-1, 0, 1}
	var hops []hop

	for _, x := range speedChange {
		for _, y := range speedChange {
			newV, err := h.v.accelerate(x, y)
			if err != nil { continue }
			newP, err := h.p.move(newV, grid)
			if err != nil { continue }

			hops = append(hops, hop{newP, newV})
		}
	}

	return hops
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
