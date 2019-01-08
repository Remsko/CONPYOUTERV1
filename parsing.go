package main

import (
	"fmt"
	"strings"
)

func parsingEntry(equation string) ([]int, []int) {

	fmt.Print("Input: ")
	fmt.Println(equation)

	leftStr, rightStr := splitEqual(equation)
	if leftStr == "" || rightStr == "" {
		return nil, nil
	}
	leftArr := parseQuadratic(leftStr)
	rightArr := parseQuadratic(rightStr)
	return leftArr, rightArr
}

func splitEqual(equation string) (string, string) {

	split := strings.Split(equation, "=")
	if split == nil || len(split) != 2 {
		return "", ""
	}
	return split[0], split[1]
}

func parseQuadratic(member string) []int {

	tokens := strings.Fields(member)

	for i := 0; i < len(tokens); i++ {

	}
	fmt.Println(tokens)
	ret := []int{0, 1, 2}

	return ret
}
