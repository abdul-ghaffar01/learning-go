package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a      int
		b      int
		result int
	}{
		{2, 3, 5},
		{2, 6, 8},
		{2, 1, 3},
	}

	for _, test := range tests {
		result := add(test.a, test.b)
		if result != test.result {
			t.Errorf("Error occured in Add(%v, %v) expected result was %v but got %v ", test.a, test.b, test.result, result)
		}
	}

}
