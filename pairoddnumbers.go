package main

import "fmt"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	// numberslice := []int{9, 3, 9, 3, 9, 7, 9}

	mapone := make(map[int]int)

	for _, a := range A {
		if mapone[a] != 0 {
			mapone[a] = mapone[a] + 1
		} else {
			mapone[a] = 1
		}
	}
	//fmt.Println(mapone)
	var count int
	for key, value := range mapone {
		if value%2 == 0 {

		} else {
			return key
		}
	}
	return count

}
func main() {
	nubmerslice := []int{9, 3, 9, 3, 9, 7, 9}
	value := Solution(nubmerslice)
	fmt.Println(value)
}
