package main

import (
	"fmt"
	"math"
)

func solution(coefs []int64) {

	var a, b, c, delta float64

	if len(coefs) > 3 {
		fmt.Println("The polynomial degree is stricly greater than 2, I can't solve.")
		return
	}
	a = float64(coefs[2])
	b = float64(coefs[1])
	c = float64(coefs[0])
	if a != float64(0) {
		delta = (b * b) - (4 * a * c)
		if delta < 0.0 {
			fmt.Printf("Discriminant is strictly negative (%.2f), the two complex solutions are:\n", delta)
			fmt.Printf("%.2f - i * %.2f / %.2f\n", -b, math.Sqrt(-delta), 2*a)
			fmt.Printf("%.2f + i * %.2f / %.2f\n", -b, math.Sqrt(-delta), 2*a)
		} else if delta == 0.0 {
			fmt.Println("The solution is:")
			root := -b / (2 * a)
			fmt.Println(root)
		} else {
			fmt.Printf("Discriminant is strictly positive (%.2f), the two solutions are:\n", delta)
			root1 := (-b + math.Sqrt(delta)) / (2 * a)
			root2 := (-b - math.Sqrt(delta)) / (2 * a)
			fmt.Println(root1)
			fmt.Println(root2)
		}
	} else if b != 0 {
		fmt.Println("The solution is:")
		if c == 0 {
			fmt.Println("0")
		} else {
			aswr := -c / b
			fmt.Println(aswr)
		}
	} else {
		if c == 0 {
			fmt.Println("All real numbers are solutions")
		} else {
			fmt.Println("There is no solution: the equation is wrong")
		}
	}
}
