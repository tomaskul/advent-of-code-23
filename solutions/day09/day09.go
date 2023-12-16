package day09

import "github.com/tomaskul/advent-of-code-23/util"

type Day09 struct {
	rows []string
}

func NewDay09Solution(sessionCookie string) *Day09 {
	rows, _ := util.GetCachedRows("https://adventofcode.com/2023/day/14/input", "09", ".txt", sessionCookie)
	return &Day09{
		rows: rows,
	}
}

func (s *Day09) PrintPart1() {
}

func (s *Day09) PrintPart2() {
}
