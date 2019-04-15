package main

import "fmt"

func main() {
	fmt.Println("Enter the number of rows ")
	var rows int
	fmt.Scanln(&rows)
	for i := 0; i <= rows-1; i++ {
		for j := 0; j <= rows-1; j++ {
			if i == 0 || i == 4 || j == 0 || j == 4 {
				fmt.Print(1)
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}
}
