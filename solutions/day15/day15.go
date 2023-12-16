package day15

import (
	"fmt"
	"strings"

	"github.com/tomaskul/advent-of-code-23/util"
)

type Day15 struct {
	rows []string
}

func NewDay15Solution(sessionCookie string) *Day15 {
	rows, _ := util.GetCachedRows("https://adventofcode.com/2023/day/15/input", "15", ".txt", sessionCookie)
	return &Day15{
		rows: rows,
	}
}

func (s *Day15) PrintPart1() {
	fmt.Println(hashSumOfTokens(strings.Split(s.rows[0], ",")))
}

func hashSumOfTokens(values []string) int {
	hashValues := make([]int, 0)
	for _, v := range values {
		hashValues = append(hashValues, hash(v))
	}
	return util.Sum(hashValues)
}

func hash(value string) int {
	var currentValue int
	for _, v := range value {
		currentValue += int(v)
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}

func (s *Day15) PrintPart2() {
}
