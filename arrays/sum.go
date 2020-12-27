package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(numbers))
}

// Sum adds together all elements of the given array
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll adds numbers in two slices and returns slice
func SumAll(slicesToSum ...[]int) (sums []int) {

	for _, numbers := range slicesToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// SumAllTails sums all elements in the slices apart from the first one
func SumAllTails(slicesToSum ...[]int) (tails []int) {

	for _, numbers := range slicesToSum {
		if len(numbers) == 0 {
			tails = append(tails, 0)
		} else {
			tail := numbers[1:]
			tails = append(tails, Sum(tail))
		}
	}
	return tails
}
