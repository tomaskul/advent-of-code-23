package day17

import "github.com/tomaskul/advent-of-code-23/util"

type Day17 struct {
	rows []string
}

func NewDay17Solution(sessionCookie string) *Day17 {
	rows, _ := util.GetCachedRows("https://adventofcode.com/2023/day/17/input", "17", ".txt", sessionCookie)
	return &Day17{
		rows: rows,
	}
}

func (s *Day17) PrintPart1() {
}

func (s *Day17) PrintPart2() {
}
