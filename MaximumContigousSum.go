package main

import "fmt"

func main() {

	numberslice := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maximumSum := MaxContigousSum(numberslice)
	fmt.Println(maximumSum)
}
func MaxContigousSum(numberslice []int) int {
	var max_ending_here int
	var max_so_far int
	for _, a := range numberslice {
		max_ending_here = max_ending_here + a
		if max_so_far < max_ending_here {
			max_so_far = max_ending_here
		}
		if max_ending_here < 0 {
			max_ending_here = 0
		}
	}
	return max_so_far

}
