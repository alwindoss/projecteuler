package problem2

import "fmt"

func Solve() {
	fibSeries := generateFibonacciNumbers()
	sum := findSumOfEvenValues(fibSeries)
	fmt.Printf("Sum: %d\n", sum)
}

func generateFibonacciNumbers() []int {
	var n []int
	first := 1
	second := 2
	n = append(n, first)
	n = append(n, second)
	for {
		newVal := first + second
		if newVal >= 4000000 {
			break
		}
		n = append(n, newVal)
		first = second
		second = newVal
	}
	return n
}

func findSumOfEvenValues(fibSeries []int) int {
	sum := 0
	for _, val := range fibSeries {
		if (val % 2) == 0 {
			sum = sum + val
		}
	}
	return sum
}
