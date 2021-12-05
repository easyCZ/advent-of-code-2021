package main

import (
	"fmt"
	"github.com/easyCZ/advent-of-code-2021/input"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := input.ReadLines(os.Stdin)
	vectors := parse(lines)

	{
		straightVectors := filterToStraightVectors(vectors)
		width, height := size(straightVectors)
		grid := NewGrid(width, height)
		grid.Plots(straightVectors)

		fmt.Println("Straight Intersections:", countNonZeroGridPositions(grid.items))
	}

	{
		width, height := size(vectors)
		grid := NewGrid(width, height)
		grid.Plots(vectors)

		fmt.Println("All intersections:", countNonZeroGridPositions(grid.items))
	}

}

func countNonZeroGridPositions(g [][]int) int {
	c := 0
	for row := range g {
		for col := range g[row] {
			if g[row][col] >= 2 {
				c++
			}
		}
	}
	return c
}

func filterToStraightVectors(vs []Vector) []Vector {
	var filtered []Vector
	for _, v := range vs {
		if v.IsStraight() {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

type Point struct {
	X int
	Y int
}

type Vector struct {
	Start Point
	End   Point
}

func (v *Vector) IsStraight() bool {
	return v.Start.X == v.End.X || v.Start.Y == v.End.Y
}

func (v *Vector) Is45Degrees() bool {
	// horizontal distance must match vertical distance, aka manhattan
	return abs(v.Start.X, v.End.X) == abs(v.Start.Y, v.End.Y)
}

//
//func (v *Vector) Points() []Point {
//	steps := abs(v.Start.X - v.End.X, v.Start.Y - v.End.Y)
//
//	for i := 0; i < steps; i++ {
//
//	}
//}

func parsePoint(s string) Point {
	tokens := strings.Split(s, ",")
	if len(tokens) != 2 {
		panic(fmt.Sprintf("failed to parse point from: %s", s))
	}

	x, err := strconv.ParseInt(tokens[0], 10, 32)
	if err != nil {
		panic(fmt.Sprintf("failed to parse X: %s", tokens[0]))
	}

	y, err := strconv.ParseInt(tokens[1], 10, 32)
	if err != nil {
		panic(fmt.Sprintf("failed to parse Y: %s", tokens[1]))
	}

	return Point{
		X: int(x),
		Y: int(y),
	}
}

func parse(lines []string) []Vector {
	var vectors []Vector
	for _, line := range lines {
		tokens := strings.Split(line, " -> ")
		if len(tokens) != 2 {
			panic(fmt.Sprintf("failed to parse ector: %s", line))
		}

		start := parsePoint(tokens[0])
		end := parsePoint(tokens[1])
		vectors = append(vectors, Vector{
			Start: start,
			End:   end,
		})
	}
	return vectors
}

type Grid struct {
	width  int
	height int

	items [][]int
}

func NewGrid(width, height int) *Grid {
	var grid [][]int
	for h := 0; h < height; h++ {
		grid = append(grid, make([]int, width))
	}

	return &Grid{
		width:  width,
		height: height,
		items:  grid,
	}
}

func (g *Grid) Plot(v Vector) {
	//(1, 1) -> (1, 4)
	// vertical
	if v.Start.X == v.End.X {

		// start is smaller
		for y := v.Start.Y; y <= v.End.Y; y++ {
			g.items[y][v.Start.X]++
		}
		// end is smaller
		for y := v.End.Y; y <= v.Start.Y; y++ {
			g.items[y][v.Start.X]++
		}
	}

	// horizontal
	if v.Start.Y == v.End.Y {
		// start is smaller
		for x := v.Start.X; x <= v.End.X; x++ {
			g.items[v.Start.Y][x]++
		}
		// end is smaller
		for x := v.End.X; x <= v.Start.X; x++ {
			g.items[v.Start.Y][x]++
		}
	}

	if v.Is45Degrees() {
		// for every step in horizotanl, we also step in diagonal. However, we can be stepping in any of the 4 directions depending on angle..
		// this is equivalent as either going top left to bottom right, or top right going to bottom left since we can reverse the vector
		xIncreases := v.Start.X < v.End.X
		yIncreases := v.Start.Y < v.End.Y
		steps := abs(v.Start.X, v.End.X)

		if xIncreases && yIncreases {
			// top left to bottom right, increase both x and y
			for i := 0; i <= steps; i++ {
				g.items[v.Start.Y+i][v.Start.X+i]++
			}
		} else if xIncreases && !yIncreases {
			// bottom left, top right
			for i := 0; i <= steps; i++ {
				g.items[v.Start.Y-i][v.Start.X+i]++
			}
		} else if !xIncreases && yIncreases {
			// top right, bottom left
			for i := 0; i <= steps; i++ {
				g.items[v.Start.Y+i][v.Start.X-i]++
			}
		} else {
			// top right, bottom left
			for i := 0; i <= steps; i++ {
				g.items[v.Start.Y-i][v.Start.X-i]++
			}
		}
	}
}

func (g *Grid) Plots(vs []Vector) {
	for _, v := range vs {
		g.Plot(v)
	}
}

func abs(x, y int) int {
	return int(math.Abs(float64(x) - float64(y)))
}

func size(vs []Vector) (int, int) {
	width := 0
	height := 0

	for _, v := range vs {
		if v.Start.X > width {
			width = v.Start.X
		}

		if v.End.X > width {
			width = v.End.X
		}

		if v.Start.Y > height {
			height = v.Start.Y
		}

		if v.End.Y > height {
			height = v.End.Y
		}
	}

	return width + 1, height + 1
}
