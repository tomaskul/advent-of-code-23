package util

import "strconv"

func ToInts(input []string) []int {
	result := make([]int, 0)
	for _, v := range input {
		integerValue, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		result = append(result, integerValue)
	}
	return result
}

func ExcludeEmptyEntries(input []string) []string {
	result := make([]string, 0)
	for _, v := range input {
		if v != "" {
			result = append(result, v)
		}
	}
	return result
}
