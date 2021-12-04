package input

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func ReadLines(r io.Reader) []string {
	scanner := bufio.NewScanner(r)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func StringsToInts(ss []string) ([]int, error) {
	var vals []int
	for _, s := range ss {
		if s == "" {
			continue
		}
		val, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("failed to parse %s: %w", s, err)
		}

		vals = append(vals, int(val))
	}
	return vals, nil
}
