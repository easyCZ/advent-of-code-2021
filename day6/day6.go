package main

import (
	"fmt"
	"github.com/easyCZ/advent-of-code-2021/input"
	"os"
	"strings"
)

func main() {
	lines := input.ReadLines(os.Stdin)
	fish := parse(lines[0])

	after80Days := simulateDaysWithCounts(counts(fish), 80)
	fmt.Println("After 80 days:", numberOfFishWithCounts(after80Days))

	countsAfter256 := simulateDaysWithCounts(counts(fish), 256)
	fmt.Println("after 256", numberOfFishWithCounts(countsAfter256))
}

func numberOfFishWithCounts(cs map[int]int) int {
	sum := 0
	for _, c := range cs {
		sum += c
	}
	return sum
}

func parse(line string) []int {
	if out, err := input.StringsToInts(strings.Split(line, ",")); err != nil {
		panic(err)
	} else {
		return out
	}

}

func counts(fish []int) map[int]int {
	cs := map[int]int{}
	for _, f := range fish {
		if _, ok := cs[f]; ok {
			cs[f]++
		} else {
			cs[f] = 1
		}
	}
	return cs
}

func simulateDayWithCounts(counts map[int]int) map[int]int {
	newCounts := map[int]int{}
	for f, c := range counts {
		if f == 0 {
			newCounts[6] += c
			newCounts[8] += c
			continue
		}

		newCounts[f-1] += c
	}
	return newCounts
}

func simulateDaysWithCounts(counts map[int]int, days int) map[int]int {
	for i := 0; i < days; i++ {
		counts = simulateDayWithCounts(counts)
	}
	return counts
}
