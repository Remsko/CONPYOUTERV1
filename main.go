package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("--- ComputerV1 ---")
	if os.Args == nil {
		fmt.Println("os.Args is nil")
		os.Exit(1)
	}
	if len(os.Args) != 2 {
		fmt.Println("usage: " + os.Args[0] + " \"<equation>\"")
		os.Exit(1)
	}
	quadra, err := parsingEntry(os.Args[1])
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	reducedCoefs := reduce(quadra.left, quadra.right)
	degree(reducedCoefs)
	solution(reducedCoefs)
}
