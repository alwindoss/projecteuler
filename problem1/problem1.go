package problem1

import "fmt"

func checkMultiple(num, mul1, mul2 int) bool {
	isMultiple := false
	if (num%mul1) == 0 || (num%mul2) == 0 {
		isMultiple = true
	}
	return isMultiple
}

func removeDuplicates(nums []int) []int {

	return nums
}

func Solve() {
	var multiples []int
	sum := 0
	for i := 1; i < 1000; i++ {
		if checkMultiple(i, 3, 5) {
			multiples = append(multiples, i)
		}
	}
	for _, val := range multiples {
		sum = sum + val
	}
	fmt.Printf("Sum: %d\n", sum)
}
