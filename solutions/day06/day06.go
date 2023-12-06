package day06

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tomaskul/advent-of-code-23/util"
)

type Day06 struct {
	rows []string
}

type race struct {
	timeMs     int
	distanceMm int
}

func (r race) noOptionsToBeat() int {
	waysToWin := []int{}
	for i := 1; i < r.timeMs; i++ {
		timeRemaining := r.timeMs - i
		distance := timeRemaining * i
		if distance > r.distanceMm {
			waysToWin = append(waysToWin, i)
		}
	}

	return len(waysToWin)
}

func NewDay06Solution(sessionCookie string) *Day06 {
	rows, _ := util.GetCachedRows("https://adventofcode.com/2023/day/6/input", "6", ".txt", sessionCookie)
	return &Day06{
		rows: rows,
	}
}

func (s *Day06) PrintPart1() {
	races := parseData(s.rows)

	result := races[0].noOptionsToBeat()
	for i := 1; i < len(races); i++ {
		result *= races[i].noOptionsToBeat()
	}

	fmt.Println(result)
}

func parseData(input []string) []race {
	timeEntries := util.ToInts(util.ExcludeEmptyEntries(strings.Split(strings.TrimPrefix(input[0], "Time: "), " ")))
	distanceEntries := util.ToInts(util.ExcludeEmptyEntries(strings.Split(strings.TrimPrefix(input[1], "Distance: "), " ")))

	races := make([]race, len(timeEntries))
	for i := 0; i < len(timeEntries); i++ {
		races[i] = race{
			timeMs:     timeEntries[i],
			distanceMm: distanceEntries[i],
		}
	}
	return races
}

func (s *Day06) PrintPart2() {
	time, _ := strconv.Atoi(strings.Join(util.ExcludeEmptyEntries(strings.Split(strings.TrimPrefix(s.rows[0], "Time: "), " ")), ""))
	distance, _ := strconv.Atoi(strings.Join(util.ExcludeEmptyEntries(strings.Split(strings.TrimPrefix(s.rows[1], "Distance: "), " ")), ""))

	race := race{
		timeMs:     time,
		distanceMm: distance,
	}

	fmt.Println(race.noOptionsToBeat())
}
