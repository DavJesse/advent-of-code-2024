package two

import "testing"

func TestFrequencyCalculator(t *testing.T) {
	subject := []int{3, 4, 2, 1, 3, 3}
	expected := map[int]int{3: 3, 4: 1, 2: 1, 1: 1}
	got := FrequencyCalculator(subject)

	for _, num := range []int{1, 2, 3, 4} {
		value1, exists := got[num]
		if exists {
			value2, _ := expected[num]

			if value1 != value2 {
				t.Errorf("Expected frequency for %d to be %d, got %d", num, value2, value1)
			}
		}
	}
}
