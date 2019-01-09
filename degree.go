package main

import "fmt"

func degree(coefs []int64) {

	var degree int64

	degree = 0
	fmt.Print("Polynomial degree: ")
	for index := len(coefs) - 1; index >= 0; index-- {
		if coefs[index] != 0 {
			degree = int64(index)
			break
		}
	}
	fmt.Println(degree)
}
