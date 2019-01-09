package main

import "fmt"

func degree(coefs []int64) {
	fmt.Print("Polynomial degree: ")
	for index := len(coefs) - 1; index >= 0; index-- {
		if coefs[index] != 0 {
			fmt.Println(index)
			return
		}
	}
}
