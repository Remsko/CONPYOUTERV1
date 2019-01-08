package main

import (
	"fmt"
	"math"
)

func solution(coefs []int) {
	var a, b, c, delta float64

	if len(coefs) > 3 {
		fmt.Println("The polynomial degree is stricly greater than 2, I can't solve.")
		return
	}
	a = float64(coefs[2])
	b = float64(coefs[1])
	c = float64(coefs[0])
	delta = (b * b) - (4 * a * c)
	if delta < 0.0 {
		fmt.Println("Discriminant is strictly negative, there is no real solutions.")
	} else if delta == 0.0 {
		fmt.Println("The solution is:")
		root := -b / (2 * a)
		fmt.Println(root)
	} else {
		fmt.Println("Discriminant is strictly positive, the two solutions are:")
		root1 := (-b + math.Sqrt(delta)) / (2 * a)
		root2 := (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Println(root1)
		fmt.Println(root2)
	}
}
