package submarine

import (
	"fmt"
	"strconv"
	"strings"
)

type Submarine struct {
	X     int
	Depth int

	Aim int
}

func (s *Submarine) Move(i Instruction) {
	switch i.Op {
	case Up:
		s.Up(i.Units)
	case Down:
		s.Down(i.Units)
	case Forward:
		s.Forward(i.Units)
	default:
		panic(fmt.Sprintf("unknown instruction: %v", i))
	}
}

func (s *Submarine) Moves(ins []Instruction) {
	for _, i := range ins {
		s.Move(i)
	}
}

func (s *Submarine) Up(i int) {
	s.Aim += -i
}

func (s *Submarine) Down(i int) {
	s.Aim += i
}

func (s *Submarine) Forward(i int) {
	s.X += i
	s.Depth += s.Aim * i
}

func (s *Submarine) String() string {
	return fmt.Sprintf("Sub{X: %d, Y: %d, Aim: %d}", s.X, s.Depth, s.Aim)
}

var (
	Up      = "up"
	Down    = "down"
	Forward = "forward"
)

type Instruction struct {
	Op    string
	Units int
}

func ParseInstruction(s string) Instruction {
	tokens := strings.Split(strings.TrimSpace(s), " ")
	op := tokens[0]
	switch tokens[0] {
	case Up, Down, Forward:
		break
	default:
		panic(fmt.Sprintf("unknown instruction %s", s))
	}

	unit, err := strconv.ParseInt(tokens[1], 10, 32)
	if err != nil {
		panic(fmt.Sprintf("unparseable unit in %s, %s", s, err))
	}

	return Instruction{
		Op:    op,
		Units: int(unit),
	}
}

func ParseInstructions(in []string) []Instruction {
	var ins []Instruction
	for _, s := range in {
		ins = append(ins, ParseInstruction(s))
	}
	return ins
}
