package util

func Sum(input []int) int {
	total := 0
	for _, value := range input {
		total += value
	}
	return total
}

func Min(input []int) int {
	min := input[0]
	for _, v := range input {
		if v < min {
			min = v
		}
	}
	return min
}
