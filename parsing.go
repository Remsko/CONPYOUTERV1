package main

import (
	"fmt"
	"strconv"
	"strings"
)

type tEqu struct {
	left  string
	right string
}

type tQuadra struct {
	left  []int64
	right []int64
}

func parsingEntry(equation string) (*tQuadra, error) {

	members, err := splitEqual(equation)
	if err != nil {
		return nil, err
	}
	leftArr, err := parseMember(members.left)
	if leftArr == nil {
		fmt.Printf("Parsing error in left member: %s\n", members.left)
		return nil, err
	}
	rightArr, err := parseMember(members.right)
	if rightArr == nil {
		fmt.Printf("Parsing error in right member: %s\n", members.right)
		return nil, err
	}
	return &tQuadra{leftArr, rightArr}, nil
}

func splitEqual(equation string) (*tEqu, error) {

	split := strings.Split(equation, "=")

	if len(split) != 2 {
		return nil, fmt.Errorf("No equal sign in equation %v", equation)
	}
	return &tEqu{split[0], split[1]}, nil
}

func parseMember(member string) ([]int64, error) {

	var length int

	tokens := strings.Fields(member)
	ret := make([]int64, 3)

	if len(tokens) == 0 {
		return nil, fmt.Errorf("Empty member")
	}
	if len(tokens) == 1 && tokens[0] == "0" {
		return ret, nil
	}
	for i := 0; i < len(tokens); i += length {

		length = 0
		if i > 0 && isSign(tokens[i]) == false {
			return nil, fmt.Errorf("Math syntax near: %v", tokens[i])
		}
		number, err := nextNumber(tokens, &length, i)
		if err != nil {
			return nil, fmt.Errorf("Expect a number but receive: %v", tokens[i])
		}
		degree, err := nextDegree(tokens, &length, i)
		if err != nil {
			return nil, fmt.Errorf("Expect a \"* X^n\" format\n")
		}
		if degree > 2 || degree < 0 {
			return nil, fmt.Errorf("ComputerV1 solve only at least equation of degree 2")
		}
		if ret[degree] != 0 {
			return nil, fmt.Errorf("Bad equation format")
		}
		ret[degree] = number
	}
	return ret, nil
}

func isSign(token string) bool {
	return token == "+" || token == "-"
}

func nextNumber(tokens []string, length *int, index int) (int64, error) {

	neg := false

	for ; index+*length < len(tokens) && isSign(tokens[index+*length]); *length++ {
		if tokens[index+*length] == "-" {
			neg = (neg == false)
		}
	}
	if index+*length >= len(tokens) {
		return 0, fmt.Errorf("Expect number")
	}
	number, err := strconv.ParseInt(tokens[index+*length], 10, 64)
	if err != nil {
		return 0, err
	}
	if neg != false {
		number = -number
	}
	return number, nil
}

func nextDegree(tokens []string, length *int, index int) (int, error) {

	err := fmt.Errorf("Wrong equation format for degrees")

	*length++
	if index+*length >= len(tokens) || tokens[index+*length] != "*" {
		return 0, err
	}
	*length++
	if index+*length >= len(tokens) {
		return 0, err
	}
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
