package day01

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/tomaskul/advent-of-code-23/util"
)

type Day01 struct {
	SessionCookie string
	rows          []string
}

type matches []string

func (m matches) First() string { return m[0] }
func (m matches) Last() string  { return m[len(m)-1] }

func (s *Day01) getData() {
	if s.rows == nil {
		s.rows = util.GetRows("https://adventofcode.com/2023/day/1/input", s.SessionCookie)
	}
}

func (s *Day01) PrintPart1() {
	s.getData()
	calibrationValues := s.getAllCalibrationValues(s.rows)
	fmt.Printf("Sum of all of the calibration values: %d\n", util.Sum(calibrationValues))
}

func (s *Day01) PrintPart2() {
	s.getData()
	calibrationValues := s.getRealCalibrationValues(s.rows)
	fmt.Printf("Sum of all of the REAL calibration values: %d\n", util.Sum(calibrationValues))
}

func (s *Day01) getAllCalibrationValues(input []string) []int {
	regex, _ := regexp.Compile("[0-9]")

	result := make([]int, len(input))
	for i, row := range input {
		matches := regex.FindAllString(row, len(row))

		digits := []string{
			matches[0],
			matches[len(matches)-1],
		}

		calibrationValue, _ := strconv.Atoi(strings.Join(digits, ""))

		result[i] = calibrationValue
	}

	return result
}

func (s *Day01) getRealCalibrationValues(input []string) []int {
	regex, _ := regexp.Compile("[0-9]")
	result := make([]int, len(input))
	for i, row := range input {
		result[i] = constructValue(regex.FindAllString(row, -1), matchSpeltNumbers(row), row)
	}

	return result
}

func matchSpeltNumbers(input string) []string {
	result := make([]string, 0)

	keys := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i := 0; i < len(input); i++ {
		for _, k := range keys {
			if len(k) > len(input[i:]) {
				continue
			}
			if strings.Contains(input[i:i+len(k)], k) {
				result = append(result, k)
				continue
			}
		}
	}

	return result
}

func constructValue(leftMatches, rightMatches matches, input string) int {
	var digitLookup = map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
	var digits []string
	if len(rightMatches) == 0 {
		digits = []string{leftMatches.First(), leftMatches.Last()}
	} else if len(leftMatches) == 0 {
		digits = []string{digitLookup[rightMatches.First()], digitLookup[rightMatches.Last()]}
	} else {
		itemOne := leftMatches.First()
		if !isLeftFirst(input, leftMatches.First(), rightMatches.First(), strings.Index) {
			itemOne = digitLookup[rightMatches.First()]
		}

		itemTwo := leftMatches.Last()
		if isLeftFirst(input, leftMatches.Last(), rightMatches.Last(), strings.LastIndex) {
			itemTwo = digitLookup[rightMatches.Last()]
		}

		digits = []string{itemOne, itemTwo}
	}

	result, _ := strconv.Atoi(strings.Join(digits, ""))

	return result
}

func isLeftFirst(search, left, right string, fn func(string, string) int) bool {
	return fn(search, left) < fn(search, right)
}
