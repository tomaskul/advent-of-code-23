package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

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

func numbersWithAdjacentSymbols(input []string) []int {
	numMatch, _ := regexp.Compile("\\d{1,7}")
	result := make([]int, 0)

	for i, row := range input {
		numbersInRow := numMatch.FindAllString(row, -1)
		if len(numbersInRow) == 0 {
			continue
		}

		rowAbove, rowBelow := "", ""
		if i-1 > -1 {
			rowAbove = input[i-1]
		}
		if i+1 < len(input) {
			rowBelow = input[i+1]
		}
		for _, number := range numbersInRow {
			if isAdjacentToSymbol(number, rowAbove, row, rowBelow) {
				value, _ := strconv.Atoi(number)
				result = append(result, value)
			}
		}
	}

	return result
}

func isAdjacentToSymbol(subject, rowAbove, subjectRow, rowBelow string) bool {
	nonFillerRegex, _ := regexp.Compile("[^0-9.]")
	rowIndex := strings.Index(subjectRow, subject)
	subjectLength := len(subject)

	leftIndex := rowIndex
	if rowIndex-1 > -1 {
		leftIndex = rowIndex - 1
		if nonFillerRegex.Match([]byte{subjectRow[leftIndex]}) {
			return true
		}
	}

	rightIndex := rowIndex + subjectLength - 1
	if rowIndex+subjectLength < len(subjectRow) {
		rightIndex = rowIndex + subjectLength
		if nonFillerRegex.Match([]byte{subjectRow[rightIndex]}) {
			return true
		}
	}

	if rowAbove != "" && nonFillerRegex.Match([]byte(rowAbove[leftIndex:rightIndex])) {
		return true
	}

	if rowBelow != "" && nonFillerRegex.Match([]byte(rowBelow[leftIndex:rightIndex])) {
		return true
	}

	return false
}
