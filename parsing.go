package main

import "fmt"

func parsingEntry(equation string) ([]int, []int) {
	left := []int{1, 2, 3}
	right := []int{4, 1, 0}

	fmt.Printf("Left part: %v\n", left)
	fmt.Printf("Right part: %v\n", right)
	return left, right
}
