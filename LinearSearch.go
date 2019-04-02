package main

import "fmt"

func main() {
	items := []int{1, 2, 57, 8, 55656, 898, 2565, 59, 65}
	fmt.Println(linearsearch(items, 565656655))
}

func linearsearch(items []int, value int) bool {
	for _, item := range items {
		if item == value {
			return true
		}
	}
	return false

}
