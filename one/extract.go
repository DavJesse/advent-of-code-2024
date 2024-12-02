package one

import (
	"sort"
)

func Extract(records []int) ([]int, []int) {
	var left, right []int

	for i := range records {
		if isEven(i) {
			left = append(left, records[i])
		} else {
			right = append(right, records[i])
		}
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	return left, right
}

func isEven(n int) bool {
	return n%2 == 0
}
