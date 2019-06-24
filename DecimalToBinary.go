package main

import (
	"fmt"
	"sort"
)

//steps
// 1.create a slice to store the binary values
// 2.loop through until decimal number is greater than zero if zero break from loop
// 3.when looping,take the remainder of decimalnumber%2 and store in slice
// 4.divide the decimal number until it is greater than zero
// 5.read the slice from reverse order

func main() {
	binaryslice := []int{}
	decimalnumber := 12
	for decimalnumber > 0 {
		binaryslice = append(binaryslice, decimalnumber%2)
		decimalnumber = decimalnumber / 2
	}
	sort.Sort(sort.Reverse(sort.IntSlice(binaryslice)))
	fmt.Println(binaryslice)
}
