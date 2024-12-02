package two

func ascending(levels []int) bool {
	return levels[0] < levels[len(levels)-1]
}

func isSafe(levels []int) bool {
	// Check for distubance in ascending or descending trends
	// Check diffence of adjacent levels
	if ascending(levels) {
		p1 := 0
		p2 := p1 + 1

		for p1 < len(levels) && p2 < len(levels) {
			if levels[p1] > levels[p2] {
				return false
			}

			difference := levels[p2] - levels[p1]
			if !(difference > 0 && difference < 4) {
				return false
			}
			p1++
			p2++
		}
	} else {
		p1 := 0
		p2 := p1 + 1

		for p1 < len(levels) && p2 < len(levels) {
			if levels[p1] < levels[p2] {
				return false
			}

			difference := levels[p1] - levels[p2]
			if !(difference > 0 && difference < 4) {
				return false
			}
			p1++
			p2++
		}
	}
	return true
}
