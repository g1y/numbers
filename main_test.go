package main

import (
	"sort"
	"testing"
)

func TestScrambledNumbersContents(t *testing.T) {
	numbers := scrambledNumbers()
	sort.Ints(numbers[:])
	for i := 0; i < len(numbers); i++ {
		if numbers[i] != i+1 {
			t.Errorf("Unexpected number. got: %d, want: %d", numbers[i], i+1)
		}
	}
}

// 0.1 ^ 20 chance that this will false positive
func TestScrambledNumbersPseudoRandomness(t *testing.T) {
	numbers := scrambledNumbers()
	if sort.IntsAreSorted(numbers[:]) {
		numbers = scrambledNumbers()
		if sort.IntsAreSorted(numbers[:]) {
			t.Errorf("List appears to be sorted. got: %p, expected unsorted", numbers[:])
		}
	}
}
