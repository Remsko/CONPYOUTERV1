package main

import (
	"fmt"
	"strconv"
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

func isSign(token string) bool {
	return token == "+" || token == "-"
}

func nextNumber(tokens []string, length *int, index int) (int, error) {

	neg := false

	for ; isSign(tokens[index+*length]) && index+*length < len(tokens); *length++ {
		if tokens[index+*length] == "-" {
			neg = (neg == false)
		}
	}
	number, err := strconv.ParseInt(tokens[index+*length], 10, 64)
	if err != nil {
		return 0, err
	}
	if neg != false {
		number = -number
	}
	return int(number), nil
}

func nextDegree(tokens []string, length *int, index int) (int, error) {

	err := fmt.Errorf("Wrong equation format for degree")
	/*
		if index+*length >= len(tokens)-2 {
			return 0, err
		}
	*/
	if tokens[index+*length] != "*" {
		return 0, err
	}
	*length++
	split := strings.Split(tokens[index+*length], "^")
	*length++
	if len(split) != 2 {
		return 0, err
	}
	if split[0] != "X" {
		return 0, err
	}
	degree, err := strconv.ParseInt(split[1], 10, 64)
	if err != nil {
		return 0, err
	}
	return int(degree), nil
}

func parseQuadratic(member string) []int {

	var length int

	tokens := strings.Fields(member)
	ret := make([]int, 3)

	for i := 0; i < len(tokens); i += length {
		length = 0
		if i > 0 && isSign(tokens[i]) == false {
			return nil
		}
		number, err := nextNumber(tokens, &length, i)
		if err != nil {
			fmt.Print(err)
			return nil
		}
		length++
		degree, err := nextDegree(tokens, &length, i)
		if err != nil {
			fmt.Print(err)
			return nil
		}
		fmt.Print("Number: ")
		fmt.Println(number)
		fmt.Print("Degree: ")
		fmt.Println(degree)
		ret[degree] = number
	}

	return ret
}
