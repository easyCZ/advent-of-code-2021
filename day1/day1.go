package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	increases := 0
	for i := 1; i < len(input); i++ {
		if input[i-1] < input[i] {
			increases++
		}
	}

	fmt.Println("Increases:", singleIncreases(input))
	fmt.Println("Window increases:", slidingWindowIncreases(input))

}

func parse(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)

	var out []int
	for scanner.Scan() {
		s := scanner.Text()
		i, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("failed to parse input '%s': %w", i, err)
		}

		out = append(out, int(i))
	}

	return out, nil
}

func singleIncreases(in []int) int {
	increases := 0
	for i := 1; i < len(in); i++ {
		if in[i-1] < in[i] {
			increases++
		}
	}
	return increases
}

func slidingWindowIncreases(in []int) int {
	var trace []int

	for i := 0; i < len(in)-2; i++ {
		sum := in[i] + in[i+1] + in[i+2]
		trace = append(trace, sum)
	}

	return singleIncreases(trace)
}
