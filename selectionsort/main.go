package main

import "fmt"

func main() {

// implementing selectionsort using nested for loops

	A := []int{5, 3, 62, 6, 33, 8}
	n := len(A)
	for i := 0; i < n; i++ {
		maxindex := i
		for j := i+1; j < n; j++ {
			if A[maxindex] > A[j] {
				maxindex = j
			}
		}
		fmt.Println("maxindex is ", maxindex)
		fmt.Println("currentindex is ", i)

		A[i], A[maxindex] = A[maxindex], A[i]
	}

	fmt.Println(A)
}
