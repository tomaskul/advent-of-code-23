package day08

import "github.com/tomaskul/advent-of-code-23/util"

type Day08 struct {
	rows []string
}

func NewDay08Solution(sessionCookie string) *Day08 {
	rows, _ := util.GetCachedRows("https://adventofcode.com/2023/day/8/input", "8", ".txt", sessionCookie)
	return &Day08{
		rows: rows,
	}
}

func (s *Day08) PrintPart1() {
}

func (s *Day08) PrintPart2() {
}
