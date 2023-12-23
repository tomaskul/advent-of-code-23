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

func BoundaryIndices(left, right, rng, minLeft, maxRight int) (int, int) {
	resultLeft, resultRight := left, right

	for i := -rng; i < left; i++ {
		if left+i > minLeft {
			resultLeft = left + i
			break
		}
	}

	for i := rng; i <= maxRight; i-- {
		if right+i <= maxRight {
			resultRight = right + i
			break
		}
	}

	return resultLeft, resultRight
}
