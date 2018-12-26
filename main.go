package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("ComputerV1:")
	if os.Args == nil {
		fmt.Println("os.Args is nil")
		os.Exit(1)
	}
	if len(os.Args) != 2 {
		fmt.Println("usage: " + os.Args[0] + " \"<equation>\"")
		os.Exit(1)
	}
	leftCoefs, rightCoefs := parsingEntry(os.Args[1])
	// reduce
	reducedCoefs := reduce(leftCoefs, rightCoefs)
	fmt.Printf("Final: %v\n", reducedCoefs)
	// degree
	degree(reducedCoefs)
	// solution
	solution(reducedCoefs)
}
