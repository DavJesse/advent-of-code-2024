package two

import "advent-of-code/one"

func FrequencyCalculator(nums []int) map[int]int {
	result := make(map[int]int)

	for _, num := range nums {
		result[num]++
	}
	return result
}

func SimilarityScoreCalculator(list []int) int {
	var frequencyScore int
	var result int
	left, right := one.Extract(list)

	// Find frequency of integers in left and right lists
	leftFreqMap := FrequencyCalculator(left)
	rightFreqMap := FrequencyCalculator(right)

	// Use keys in left list to find out if they exist in right list
	// Calculate frequency score and accumulate in result
	for key, valueLeft := range leftFreqMap {
		valueRight, exists := rightFreqMap[key]

		if exists {
			frequencyScore = key * valueRight * valueLeft
			result += frequencyScore
		}
	}

	return result
}
