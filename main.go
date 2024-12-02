package main

import (
	"advent-of-code/one"
	"advent-of-code/two"

	"fmt"
)

func main() {
	result := one.DayOnePart1Solution()
	fmt.Println("Day One Part 1 Solution")
	fmt.Println("Sum of differences:", result)
	fmt.Println()

	result = one.DayOnePart2Solution()
	fmt.Println("Day One Part 2 Solution")
	fmt.Println("Similarity score:", result)
	fmt.Println()

	result = two.DayTwoSolution()
	fmt.Println("Day Two Solution: 'Red Nose Report'")
	fmt.Println("Number of safe reports:", result)
	fmt.Println()

}
