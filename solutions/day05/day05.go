package day05

import (
	"fmt"
	"strings"

	"github.com/tomaskul/advent-of-code-23/util"
)

type Day05 struct {
	SessionCookie string
	rows          []string
}

type data struct {
	seeds []int
	paths []lookup
}

type lookup map[int]idRangeDef

type idRangeDef struct {
	srcRangeStart int
	rangeLen      int
}

func (l lookup) getSrc2DstMap(destinationId int) map[int]int {
	result := map[int]int{}
	for i := 0; i < l[destinationId].rangeLen; i++ {
		result[l[destinationId].srcRangeStart+i] = destinationId + i
	}
	return result
}

func (l lookup) getDestinationId(sourceId int) int {
	result := -1
	for destId, rng := range l {
		// if within lookup item range, find actual destinationId, otherwise try another lookup item.
		if sourceId >= rng.srcRangeStart && sourceId <= rng.srcRangeStart+(rng.rangeLen-1) {
			diff := sourceId - rng.srcRangeStart
			return destId + diff
		}
	}

	if result == -1 {
		return sourceId
	}
	return result
}

func (s *Day05) getData() {
	if s.rows == nil {
		s.rows, _ = util.GetCachedRows("https://adventofcode.com/2023/day/5/input", "5", ".txt", s.SessionCookie)
	}
}

func (s *Day05) PrintPart1() {
	s.getData()

	result := traverse(parseData(s.rows))

	fmt.Println(util.Min(result))
}

func parseData(rows []string) data {
	result := data{
		seeds: util.ToInts(strings.Split(strings.TrimPrefix(rows[0], "seeds: "), " ")),
		paths: make([]lookup, 0),
	}

	for i := 3; i < len(rows); i++ {
		nextIndex, pathMap := parseMap(rows[i:])
		if len(pathMap) == 0 {
			break
		}

		result.paths = append(result.paths, pathMap)
		if nextIndex == 0 || nextIndex == -1 {
			fmt.Printf("debug: hello something happened!")
			break
		} else {
			i = nextIndex + i + 1
		}
	}

	return result
}

func parseMap(rows []string) (int, map[int]idRangeDef) {
	result := make(map[int]idRangeDef)
	idx := -1
	for i, row := range rows {
		idx = i
		if row == "" {
			break
		}
		values := util.ToInts(strings.Split(row, " "))
		if len(values) != 3 {
			fmt.Printf("err: %q doesn't contain expected number of numbers\n", row)
			continue
		}
		result[values[0]] = idRangeDef{
			srcRangeStart: values[1],
			rangeLen:      values[2],
		}
	}
	return idx, result
}

func traverse(data data) []int {
	result := make([]int, len(data.seeds))
	for idx, seedId := range data.seeds {
		srcId := seedId
		for _, path := range data.paths {
			srcId = path.getDestinationId(srcId)
		}
		result[idx] = srcId
	}
	return result
}

func (s *Day05) PrintPart2() {
	s.getData()
}
