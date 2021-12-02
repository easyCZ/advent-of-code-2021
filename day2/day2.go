package main

import (
	"bufio"
	"fmt"
	"github.com/easyCZ/advent-of-code-2021/submarine"
	"io"
	"os"
)

func main() {
	in := parse(os.Stdin)
	instructions := submarine.ParseInstructions(in)

	sub := submarine.Submarine{}
	sub.Moves(instructions)

	fmt.Println(sub)
	fmt.Println("X * Depth = ", sub.X*sub.Depth())
}

func parse(r io.Reader) []string {
	scanner := bufio.NewScanner(r)

	var out []string
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	return out
}
