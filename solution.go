package main

import "fmt"

func solution(coefs []int) {
	if len(coefs) > 3 {
		fmt.Println("The polynomial degree is stricly greater than 2, I can't solve.")
		return
	}

	fmt.Println("Discriminant is strictly positive, the two solutions are:")
	fmt.Println("The solution is:")
}
