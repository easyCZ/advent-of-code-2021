package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/easyCZ/advent-of-code-2021/input"
)

func main() {
	lines := input.ReadLines(os.Stdin)

	gamma, epsilon := rates(lines)
	fmt.Println("Gamma:", gamma)
	fmt.Println("Epsilon:", epsilon)
	fmt.Println("Gamma * Epsilon:", gamma*epsilon)

	ox := oxygen(lines)
	scrub := scrubber(lines)
	fmt.Println("Oxygen:", ox)
	fmt.Println("Scrubber:", scrub)
	fmt.Println("Life support:", ox*scrub)
}

func rates(lines []string) (int, int) {
	// most common in corresponding position
	ones := make([]int, len(lines[0]))
	zeroes := make([]int, len(lines[0]))
	for i := 0; i < len(lines[0]); i++ {
		for _, line := range lines {
			if line[i] == '0' {
				zeroes[i] += 1
			}
			if line[i] == '1' {
				ones[i] += 1
			}
		}
	}

	fmt.Println(ones)
	fmt.Println(zeroes)

	gamma := ""
	for i := 0; i < len(lines[0]); i++ {
		if ones[i] > zeroes[i] {
			gamma += "1"
		} else {
			gamma += "0"
		}
	}

	// XOR would probs work the same but this is easier (copy/paste)

	epsilon := ""
	for i := 0; i < len(lines[0]); i++ {
		if ones[i] > zeroes[i] {
			epsilon += "0"
		} else {
			epsilon += "1"
		}
	}

	return toDecimal(gamma), toDecimal(epsilon)
}

func toDecimal(s string) int {
	parsed, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(fmt.Sprintf("failed to parse %s, %s", s, err))
	}

	return int(parsed)
}

func countValuesAtPos(lines []string, position int) (zero int, one int) {
	var vals []byte

	for _, line := range lines {
		vals = append(vals, line[position])
	}

	for _, val := range vals {
		if val == '0' {
			zero++
		} else {
			one++
		}
	}

	return zero, one
}
func mostCommonValue(lines []string, position int, equalVal byte) byte {
	zero, one := countValuesAtPos(lines, position)

	if zero == one {
		return equalVal
	}
	if zero > one {
		return '0'
	}
	return '1'
}

func leastCommonValue(lines []string, position int, equalVal byte) byte {
	zero, one := countValuesAtPos(lines, position)

	if zero == one {
		return equalVal
	}
	if zero > one {
		return '1'
	}
	return '0'
}

func filterByPositionEqual(lines []string, pos int, value byte) []string {
	var results []string
	for _, line := range lines {
		if line[pos] == value {
			results = append(results, line)
		}
	}

	return results
}

func oxygen(lines []string) int {
	for i := 0; i < len(lines[0]); i++ {
		if len(lines) == 0 {
			panic("no lines")
		}
		if len(lines) == 1 {
			return toDecimal(lines[0])
		}

		filter := mostCommonValue(lines, i, '1')
		lines = filterByPositionEqual(lines, i, filter)
	}

	if len(lines) == 1 {
		return toDecimal(lines[0])
	}
	panic("no solution found")
}

func scrubber(lines []string) int {
	for i := 0; i < len(lines[0]); i++ {
		if len(lines) == 0 {
			panic("no lines")
		}
		if len(lines) == 1 {
			return toDecimal(lines[0])
		}

		filter := leastCommonValue(lines, i, '0')
		lines = filterByPositionEqual(lines, i, filter)
	}

	if len(lines) == 1 {
		return toDecimal(lines[0])
	}
	panic("no solution found")
}
