package main

import (
	"fmt"
	"github.com/easyCZ/advent-of-code-2021/input"
	"math"
	"os"
	"strings"
)

func main() {
	lines := input.ReadLines(os.Stdin)
	crabs, err := input.StringsToInts(strings.Split(lines[0], ","))
	if err != nil {
		panic(err)
	}

	fmt.Println("Fuel constant cost:", bruteForce(crabs, constantCost))
	fmt.Println("Fuel linear cost:", bruteForce(crabs, linearCost))
}

func bruteForce(crabs []int, costFn func(int) int) int {
	min, max := crabs[0], crabs[0]
	for _, c := range crabs {
		if c < min {
			min = c
		}
		if c > max {
			max = c
		}
	}

	minFuel := math.MaxInt
	for i := min; i <= max; i++ {
		fuel := 0
		for _, c := range crabs {
			fuel += costFn(int(math.Abs(float64(i) - float64(c))))
		}

		if fuel < minFuel {
			minFuel = fuel
		}
	}

	return minFuel
}

func constantCost(n int) int {
	return n
}

func linearCost(n int) int {
	// aka sum of a series
	return n * (n+1)/2
}
