package two

import (
	"strconv"
	"strings"
)

func toInteger(reports []string) [][]int {
	var result [][]int
	var line []int

	for _, report := range reports {
		nums := strings.Split(report, "")

		for _, num := range nums {
			value, _ := strconv.Atoi(num)
			line = append(line, value)
		}

		result = append(result, line)
		line = nil
	}
	return result
}
