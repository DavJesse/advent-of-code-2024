package two

func FrequencyCalculator(nums []int) map[int]int {
	result := make(map[int]int)

	for _, num := range nums {
		result[num]++
	}
	return result
}
