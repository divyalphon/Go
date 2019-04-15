package main

import "fmt"

func main() {
	var sizeofArray int

	fmt.Println("Enter the size of an array")
	fmt.Scanln(&sizeofArray)
	var arrayone [100]int
	fmt.Println("Enter the elements of array")

	for i := 0; i < sizeofArray; i++ {
		fmt.Scanln(&arrayone[i])
	}

	sum := findingsum(arrayone)
	average := sum / sizeofArray
	fmt.Println(" average   ", average)

}

func findingsum(arrayone [100]int) int {
	sum := 0

	for i := 0; i < len(arrayone); i++ {
		sum = sum + arrayone[i]
	}
	return sum
}
