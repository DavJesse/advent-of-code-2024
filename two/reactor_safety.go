package two

func ascending(levels []int) bool {
	return levels[0] < levels[len(levels)-1]
}
