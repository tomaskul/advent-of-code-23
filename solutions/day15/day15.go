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
	tokens := strings.Split(s.rows[0], ",")

	//strings.Split("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7", ",")
	out := hashSumOfTokens(tokens)
	fmt.Printf("%d\n", out)
}

func hashSumOfTokens(values []string) int {
	//
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
