package day03

import (
	"fmt"

	"github.com/tomaskul/advent-of-code-23/util"
)

type Day03 struct {
	SessionCookie string
	rows          []string
}

func (s *Day03) getData() {
	if s.rows == nil {
		s.rows = util.GetRows("https://adventofcode.com/2023/day/3/input", s.SessionCookie)
	}
}

func (s *Day03) PrintPart1() {
	s.getData()

	fmt.Println("WIP")
}

func (s *Day03) PrintPart2() {
	s.getData()

	fmt.Println("WIP")
}

func parseInputData(input []string) int { return 0 }
