package main

import (
	"fmt"
	"github.com/easyCZ/advent-of-code-2021/input"
	"os"
	"strings"
)

func main() {
	lines := input.ReadLines(os.Stdin)
	draws, boards := parse(lines)

	for _, b := range boards {
		fmt.Println(b)
		fmt.Println("")
	}

	lastDraw, winningBoard := Play(draws, boards)
	fmt.Println("Last draw:", lastDraw)
	fmt.Println("Board:", winningBoard)
	sumUnmarked := winningBoard.SumUnmarked()
	fmt.Println("Sum unmarked:", sumUnmarked)
	fmt.Println("Sum unmarked * Last draw:", lastDraw*sumUnmarked)
}

func Play(draws []int, boards []*BingoBoard) (int, *BingoBoard) {
	for _, draw := range draws {
		// mark each board, then check who won
		for _, board := range boards {
			board.Mark(draw)

			if board.Wins() {
				return draw, board
			}
		}
	}

	return -1, nil
}

func NewBingoBoard(g [][]int) *BingoBoard {
	return &BingoBoard{
		Grid:   g,
		Marked: make(map[int]struct{}),
	}
}

type BingoBoard struct {
	Grid   [][]int
	Marked map[int]struct{}
}

func (b BingoBoard) Wins() bool {
	// horizontal
	for i := 0; i < 5; i++ {
		row := selectRow(b, i)
		if allMarked(row, b.Marked) {
			return true
		}
		col := selectColumn(b, i)
		if allMarked(col, b.Marked) {
			return true
		}
	}

	return false
}

func (b BingoBoard) SumUnmarked() int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(b.Grid)
			val := b.Grid[i][j]
			if _, ok := b.Marked[val]; ok {
				continue
			}
			sum += val
		}
	}
	return sum
}

func (b *BingoBoard) Mark(i int) {
	b.Marked[i] = struct{}{}
}

func (b *BingoBoard) String() string {
	var lines []string
	for _, row := range b.Grid {
		lines = append(lines, fmt.Sprintf("%v", row))
	}

	return strings.Join(lines, "\n")
}

func selectRow(board BingoBoard, i int) []int {
	return board.Grid[i]
}

func selectColumn(board BingoBoard, i int) []int {
	var nums []int
	for row := 0; row < 5; row++ {
		nums = append(nums, board.Grid[row][i])
	}
	return nums
}

func allMarked(nums []int, marked map[int]struct{}) bool {
	for _, num := range nums {
		if _, ok := marked[num]; !ok {
			return false
		}
	}
	return true
}

func parse(lines []string) ([]int, []*BingoBoard) {
	draws, err := input.StringsToInts(strings.Split(lines[0], ","))
	if err != nil {
		panic(err)
	}

	var boards []*BingoBoard

	var boardRows []string
	for _, line := range lines[1:] {
		if strings.TrimSpace(line) == "" {
			// baord delimiter, finalize existing board and continue
			if len(boardRows) == 0 {
				continue
			}

			boards = append(boards, parseBoard(boardRows))
			boardRows = nil
			continue
		}

		boardRows = append(boardRows, line)
	}

	return draws, boards
}

func parseBoard(rows []string) *BingoBoard {
	var parsed [][]int
	for _, row := range rows {
		tokens := strings.Split(row, " ")
		p, err := input.StringsToInts(tokens)
		if err != nil {
			panic(err)
		}
		parsed = append(parsed, p)
	}

	return NewBingoBoard(parsed)
}
