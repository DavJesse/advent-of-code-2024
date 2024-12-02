package one

func DayOnePart1Solution() int {
	left, right := Extract(Records)
	result := CalculateSumOfDifferences(left, right)
	return result
}
