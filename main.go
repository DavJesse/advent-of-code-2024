package main

import (
	"advent-of-code/one"
	"advent-of-code/two"
	"fmt"
)

func main() {
	result := one.DayOneSolution()
	fmt.Println("Day One Solution")
	fmt.Println("Sum of differences:", result)
	fmt.Println()

	result = two.DayTwoSolution()
	fmt.Println("Day Two Solution")
	fmt.Println("Similarity score:", result)
	fmt.Println()

}
