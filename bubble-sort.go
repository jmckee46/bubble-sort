package main

import "fmt"

func countSwaps(a []int32) {
	n := int32(len(a))
	var temp int32
	var swaps int32

	for i := int32(0); i < n; i++ {

		for j := int32(0); j < n-1; j++ {
			// Swap adjacent elements if they are in decreasing order
			if a[j] > a[j+1] {
				temp = a[j]
				a[j] = a[j+1]
				a[j+1] = temp
				swaps++
			}
		}
	}

	fmt.Printf("Array is sorted in %d swaps.\n", swaps)
	fmt.Println("First Element:", a[0])
	fmt.Println("Last Element:", a[n-1])
}

func main() {}
