package util

func Sum(values []int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func Min(values []int) int {
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

func All(values []int, condition func(int) bool) bool {
	for _, v := range values {
		if !condition(v) {
			return false
		}
	}
	return true
}
