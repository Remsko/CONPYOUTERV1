package main

import "fmt"

func reduce(left []int, right []int) []int {
	var value int

	fmt.Print("Reduced form:")
	length := len(left)
	coefficients := make([]int, length)
	count := 0
	for index := 0; index < length; index++ {
		value = left[index] - right[index]
		coefficients[index] = value

		if value != 0 {
			if value < 0 {
				fmt.Printf(" - %d * X^%d", -value, index)
			} else if count == 0 {
				fmt.Printf(" %d * X^%d", value, index)
			} else {
				fmt.Printf(" + %d * X^%d", value, index)
			}
			count++
		}
	}
	if count == 0 {
		fmt.Print(" 0")
	}
	fmt.Println(" = 0")
	return coefficients
}
