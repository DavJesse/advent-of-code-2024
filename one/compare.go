package one

import "math"

func CalculateSumOfDifferences(left, right []int) int {
	var result, p1, p2 int

	for p1 < len(left) && p2 < len(right) {
		result += int(math.Abs(float64(left[p1] - right[p2])))
		p1++
		p2++
	}
	return result
}
