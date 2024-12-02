package one

func DayOneSolution() int {
	left, right := Extract(Records)
	result := CalculateSumOfDifferences(left, right)
	return result
}
