package solutions

import (
	"fmt"

	"github.com/tomaskul/advent-of-code-23/solutions/day01"
	"github.com/tomaskul/advent-of-code-23/solutions/day02"
	"github.com/tomaskul/advent-of-code-23/solutions/day04"
	"github.com/tomaskul/advent-of-code-23/solutions/day05"
	"github.com/tomaskul/advent-of-code-23/solutions/day06"
	"github.com/tomaskul/advent-of-code-23/solutions/day07"
	"github.com/tomaskul/advent-of-code-23/solutions/day14"
	"github.com/tomaskul/advent-of-code-23/solutions/day15"
	"github.com/tomaskul/advent-of-code-23/solutions/day17"
)

type Solution interface {
	PrintPart1()
	PrintPart2()
}

type SolutionRegistry struct {
	registry map[int]func(string) Solution
}

func NewSolutionRegistry() *SolutionRegistry {
	return &SolutionRegistry{
		registry: map[int]func(string) Solution{
			1: func(s string) Solution { return day01.NewDay01Solution(s) },
			2: func(s string) Solution { return day02.NewDay02Solution(s) },

			4: func(s string) Solution { return day04.NewDay04Solution(s) },
			5: func(s string) Solution { return day05.NewDay05Solution(s) },
			6: func(s string) Solution { return day06.NewDay06Solution(s) },
			7: func(s string) Solution { return day07.NewDay07Solution(s) },

			14: func(s string) Solution { return day14.NewDay14Solution(s) },
			15: func(s string) Solution { return day15.NewDay15Solution(s) },

			17: func(s string) Solution { return day17.NewDay17Solution(s) },
		},
	}
}

func (r *SolutionRegistry) Get(sessionCookie string, day int) (Solution, error) {
	if day < 1 || day > 25 {
		return nil, fmt.Errorf("invalid day number: %d", day)
	}

	if createSolutionFunc, ok := r.registry[day]; ok {
		return createSolutionFunc(sessionCookie), nil
	}
	return nil, fmt.Errorf("day %d hasn't been solved by me", day)
}
