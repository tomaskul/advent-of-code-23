package main

import (
	"flag"
	"fmt"
	"os"

	s "github.com/tomaskul/advent-of-code-23/solutions"
)

func main() {
	var sessionCookie string
	var day int
	flag.StringVar(&sessionCookie, "s", "", "Session cookie to auth & retrieve user specific problem.")
	flag.IntVar(&day, "day", 0, "Day number for which to run solution for.")
	flag.Parse()

	if sessionCookie == "" {
		fmt.Println("Invalid session cookie supplied")
		os.Exit(-1)
	}

	registry := s.NewSolutionRegistry()

	printBanner()
	solution, err := registry.Get(sessionCookie, day)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("### [Day %d solution Part 1] ###\n", day)
	solution.PrintPart1()

	fmt.Printf("\n\n### [Day %d solution Part 2] ###\n", day)
	solution.PrintPart2()
}

func printBanner() {
	fmt.Println("  *\t\t\t      *")
	fmt.Println(" / \\       +++ ~~~ +++       / \\")
	fmt.Printf("/   \\  Advent of Code 2023  /   \\\n")
	fmt.Printf("''I''      +++ ~~~ +++      ''I''\n\n")
}
