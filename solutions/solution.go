package solutions

import (
	"fmt"
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
			//1: &day01.Day01{SessionCookie: sessionCookie},
			//2: ...
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
