package main

import "fmt"

func main() {

// implementing bubbleSort using nested for loop

	arr := []int{5, 3, 62, 6, 33, 8}
	n := len(A)

	// the 'len(arr)-1' is so that the loop doesnt go to the end of the
	// array as the last element will be sorted
	for i := 0; i < len(arr)-1; i++ {
		/*
			the 'j<len(arr)-i-1' is so that the inner loop sorts
			thru the unsorted parts, saves time as its unnecessary to
			sort thru sorted parts of the slice
		*/
		for j := 0; j < len(arr)-i-1; j++ {
			fmt.Println("this the len of arr", len(arr)-i-1)

			if arr[j] > arr[j+1] {
				// The current index and the one after it
				// are compared to see which is bigger and then swapped
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
