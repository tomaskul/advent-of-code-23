package solutions

import (
	"fmt"

	"github.com/tomaskul/advent-of-code-23/solutions/day01"
	"github.com/tomaskul/advent-of-code-23/solutions/day02"
	"github.com/tomaskul/advent-of-code-23/solutions/day04"
	"github.com/tomaskul/advent-of-code-23/solutions/day05"
)

type Solution interface {
	PrintPart1()
	PrintPart2()
}

type SolutionRegistry struct {
	registry map[int]Solution
}

func NewSolutionRegistry(sessionCookie string) *SolutionRegistry {
	return &SolutionRegistry{
		registry: map[int]Solution{
			1: day01.NewDay01Solution(sessionCookie),
			2: day02.NewDay02Solution(sessionCookie),

			4: day04.NewDay04Solution(sessionCookie),
			5: day05.NewDay05Solution(sessionCookie),
		},
	}
}

func (r *SolutionRegistry) Get(day int) (Solution, error) {
	if day < 1 || day > 25 {
		return nil, fmt.Errorf("invalid day number: %d", day)
	}

	if solution, ok := r.registry[day]; ok {
		return solution, nil
	}
	return nil, fmt.Errorf("day %d hasn't been solved by me", day)
}
