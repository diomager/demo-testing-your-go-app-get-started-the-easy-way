package main

import "fmt"

// Avg calculates the average of all integers
func Avg(nos ...int) int {
	sum := 0
	for _, n := range nos {
		sum += n
	}
	if sum == 0 {
		return 0
	}
	return sum / len(nos)
}

func main() {
	numbers := []int{1, 4, 5, 6}

	fmt.Println(Avg(numbers...))
}
