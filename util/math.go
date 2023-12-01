package util

func Sum(input []int) int {
	total := 0
	for _, value := range input {
		total += value
	}
	return total
}
