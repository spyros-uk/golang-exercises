package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(5, 5)
	if total != 10 {
		t.Errorf("Add - actual: %d, expected: %d", total, 10)
	}
}

func TestMultiply(t *testing.T) {
	type testCase struct {
		a        int
		b        int
		expected int
	}

	testCases := []testCase{
		{1, 5, 5},
		{3, 5, 15},
		{2, 10, 20},
	}

	for id, tc := range testCases {
		total := Multiply(tc.a, tc.b)
		if total != tc.expected {
			t.Errorf("Multiply case %d - Multiplying %d with %d got %d, expected: %d", id, tc.a, tc.b, total, tc.expected)
		}
	}
}
