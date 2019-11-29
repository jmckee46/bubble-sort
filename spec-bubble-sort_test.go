package main

import "testing"

func TestBubbleSort1(t *testing.T) {
	a := []int32{int32(6), int32(4), int32(1)}

	countSwaps(a)

	// expected output:
	// Array is sorted in 3 swaps
	// First Element: 1
	// Last Element: 6
}

func TestBubbleSort2(t *testing.T) {
	a := []int32{int32(1), int32(2), int32(3)}

	countSwaps(a)

	// expected output:
	// Array is sorted in 0 swaps
	// First Element: 1
	// Last Element: 3
}
