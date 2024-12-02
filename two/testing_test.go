package two

import "testing"

func TestToInteger(t *testing.T) {
	subject := []string{
		"1 2 3 4 5 6 7 8 9 10",
		"11 12 13 14 15 16 17 18 19 20",
	}

	expected := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
	}

	got := toInteger(subject)

	if len(got) != len(expected) {
		t.Errorf("Expected size %d, got %d", len(expected), len(got))
		t.Errorf("TestToInteger Failed!")
	}

	var r1, r2, p1, p2 int
	for r1 < len(got) && r2 < len(expected) {
		if len(got[r1]) != len(expected[r2]) {
			t.Errorf("Expected %v row at index %d to have a size of %d, got %v of size %d", expected[r2], r2, len(expected[r2]), got[r1], len(got[r1]))
			t.Errorf("TestToInteger Failed!")
		}
		for p1 < len(got[r1]) && p2 < len(expected[r2]) {
			if got[r1][p1] != expected[r2][p2] {
				t.Errorf("Expected %d at index %d, %d at index %d to be equal", expected[r2][p2], p2, got[r1][p1], p1)
				t.Errorf("TestToInteger Failed!")
			}
			p1++
			p2++
		}
		r1++
		r2++
	}

}

func TestAscending(t *testing.T) {
	subject := [][]int{
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
	}

	expected := []bool{
		true,
		false,
	}

	var p1, p2 int

	for p1 < len(subject) && p2 < len(expected) {
		if ascending(subject[p1]) != expected[p2] {
			t.Errorf("Expected %t at index %d to be equal to %t", expected[p2], p2, ascending(subject[p1]))
			t.Errorf("TestAscending Failed!")
		}
		p1++
		p2++
	}
}

func TestIsValid(t *testing.T) {
	subject := [][]int{
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{1, 2, 2, 4, 6},
		{1, 2, 8, 9, 10},
		{10, 6, 3, 2, 1},
	}

	expected := []bool{
		true,
		true,
		false,
		false,
		false,
	}

	var p1, p2 int

	for p1 < len(subject) && p2 < len(expected) {
		if isValid(subject[p1]) != expected[p2] {
			t.Errorf("Expected %t at index %d to be equal to %t", expected[p2], p2, ascending(subject[p1]))
			t.Errorf("TestIsValid Failed!")
		}
		p1++
		p2++
	}
}
