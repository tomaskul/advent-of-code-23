package day07

import "github.com/tomaskul/advent-of-code-23/util"

type Day07 struct {
	rows []string
}

func NewDay07Solution(sessionCookie string) *Day07 {
	rows, _ := util.GetCachedRows("https://adventofcode.com/2023/day/7/input", "7", ".txt", sessionCookie)
	return &Day07{
		rows: rows,
	}
}

func (s *Day07) PrintPart1() {
}

func (s *Day07) PrintPart2() {
}
