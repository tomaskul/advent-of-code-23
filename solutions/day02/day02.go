package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tomaskul/advent-of-code-23/util"
)

type Day02 struct {
	SessionCookie string
	rows          []string
}

type rgb struct {
	r int
	g int
	b int
}

type set []rgb
type game struct {
	id   int
	sets set
}

func (s *Day02) getData() {
	if s.rows == nil {
		s.rows = util.GetRows("https://adventofcode.com/2023/day/2/input", s.SessionCookie)
	}
}

func (s *Day02) PrintPart1() {
	s.getData()

	fmt.Println(util.Sum(countPossibleGames(rgb{r: 12, g: 13, b: 14}, parseInputData(s.rows))))
}

func (s *Day02) PrintPart2() {
	s.getData()
}

func countPossibleGames(expectation rgb, games []game) []int {

	result := []int{}
	for _, game := range games {
		exceedsCount := false
		for _, rgb := range game.sets {
			if rgb.r > expectation.r || rgb.g > expectation.g || rgb.b > expectation.b {
				exceedsCount = true
				break
			}
		}
		if !exceedsCount {
			result = append(result, game.id)
		}
	}

	return result
}

func parseInputData(rows []string) []game {
	gamesPlayed := make([]game, len(rows))
	for i, row := range rows {
		game := game{}

		tokens := strings.Split(row, ":")
		game.id, _ = strconv.Atoi(strings.TrimPrefix(tokens[0], "Game "))
		game.sets = parseGameSets(strings.Split(tokens[1], ";"))

		gamesPlayed[i] = game
	}

	//fmt.Printf("debug: [:5] games: %+v\n\n", gamesPlayed[:5])
	//fmt.Printf("debug: [95:] games: %+v\n\n", gamesPlayed[95:])
	return gamesPlayed
}

func parseGameSets(setTokens []string) set {
	result := make(set, len(setTokens))
	for i, revealedSet := range setTokens {
		result[i] = parseRgb(strings.Split(revealedSet, ","))
	}

	return result
}

func parseRgb(cubesDrawn []string) rgb {
	result := rgb{}
	for _, token := range cubesDrawn {
		if strings.Contains(token, "red") {
			result.r, _ = strconv.Atoi(strings.TrimSpace(strings.TrimSuffix(token, " red")))
		} else if strings.Contains(token, "green") {
			result.g, _ = strconv.Atoi(strings.TrimSpace(strings.TrimSuffix(token, " green")))
		} else if strings.Contains(token, "blue") {
			result.b, _ = strconv.Atoi(strings.TrimSpace(strings.TrimSuffix(token, " blue")))
		}
	}
	return result
}
