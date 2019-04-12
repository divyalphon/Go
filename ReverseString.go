package main

import "fmt"

func main() {
	fmt.Println("Enter the string to reverse")
	var strone string
	fmt.Scanln(&strone)
	var runes = []rune(strone)

	for i := len(runes) - 1; i >= 0; i-- {
		fmt.Print(string(runes[i]))
	}
	fmt.Println()
	//second method
	fmt.Println(reverse(strone))
}

func reverse(strone string) string {
	aa := []rune(strone)
	for i, j := 0, len(aa)-1; i < j; i, j = i+1, j-1 {
		aa[i], aa[j] = aa[j], aa[i]
	}
	return string(aa)
}
