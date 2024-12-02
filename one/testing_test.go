package one

import (
	"testing"
)

func TestExtract(t *testing.T) {
	subject := []int{9, 10, 7, 8, 5, 6, 3, 4, 1, 2}
	expected1, expected2 := []int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8, 10}
	got1, got2 := Extract(subject)

	var p1, p2 int
	if len(got1) != len(expected1) || len(got2) != len(expected2) {
		if len(got1) != len(expected1) {
			t.Errorf("Expected result of size %d, got size of %d", len(expected1), got1)
		} else {
			t.Errorf("Expected result of size %d, got size of %d", len(expected2), got2)
		}
	}

	for p1 < len(got1) && p2 < len(got2) {
		if got1[p1] != expected1[p1] || got2[p2] != expected2[p2] {
			if got1[p1] != expected1[p1] {
				t.Errorf("Expected %d at position %d, got %d", expected1[p1], p1, got1[p1])

			} else {
				t.Errorf("Expected %d at position %d, got %d", expected2[p2], p2, got2[p2])
			}
		}
		p1++
		p2++
	}
}

func TestCalculateSumOfDifference(t *testing.T) {
	subject1, subject2 := []int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8, 10}
	got := CalculateSumOfDifferences(subject1, subject2)
	expected := 5

	if got != expected {
		t.Errorf("Expected sum of differences to be %d, got %d", expected, got)
	}
}
