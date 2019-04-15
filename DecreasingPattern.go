package main

import "fmt"

func main() {
	fmt.Println("Enter the rows from where it has to be decreased")
	var rows int
	fmt.Scanln(&rows)
	for i := rows - 1; i >= 0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}
