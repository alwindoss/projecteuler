package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alwindoss/projecteuler/problem1"
	"github.com/alwindoss/projecteuler/problem2"
)

func main() {
	progPtr := flag.Int("prog", 0, "-prog 1")
	flag.Parse()
	if *progPtr == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	problemNum := *progPtr
	switch problemNum {
	case 1:
		fmt.Println("In case 1")
		problem1.Solve()
	case 2:
		fmt.Println("In case 2")
		problem2.Solve()
	default:
		fmt.Println("No Problems available yet")
	}
}
