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
}

func (s *Day01) PrintPart1() {
	rows := util.GetRows("https://adventofcode.com/2023/day/1/input", s.SessionCookie)
	calibrationValues, err := s.getAllCalibrationValues(rows)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	total := 0
	for _, value := range calibrationValues {
		total += value
	}

	fmt.Printf("Sum of all of the calibration values: %d\n", total)
}

func (s *Day01) PrintPart2() {}

func (s *Day01) getAllCalibrationValues(input []string) ([]int, error) {
	regex, err := regexp.Compile("[0-9]")
	if err != nil {
		return nil, fmt.Errorf("regexp error: %v\n", err)
	}

	result := make([]int, len(input))
	for i, row := range input {
		matches := regex.FindAllString(row, len(row))
		if matches == nil {
			fmt.Errorf("matches nil: %v", matches)
			continue
		}

		digits := []string{
			matches[0],
			matches[len(matches)-1],
		}

		calibrationValue, err := strconv.Atoi(strings.Join(digits, ""))
		if err != nil {
			fmt.Errorf("couldn't turn matches into value: %v", err)
			continue
		}

		fmt.Printf("value: %d\n", calibrationValue)
		result[i] = calibrationValue
	}

	return result, nil
}
