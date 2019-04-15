package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter the number")
	fmt.Scanln(&number)

	Reverse(number)
}

func Reverse(number int) {

	var coynumber = number
	var reverse int
	var remainder int
	for {
		remainder = coynumber % 10
		reverse = reverse*10 + remainder
		coynumber = coynumber / 10
		if coynumber == 0 {
			break
		}
	}
	if reverse == number {
		fmt.Println("Number is palindrome")
	} else {
		fmt.Println("Number is not palindrome")
	}
}
