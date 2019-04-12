package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Enterh your mobile number")
	var number string
	fmt.Scanln(&number)
	if (numberornot(number)) || len(number) == 10 {
		fmt.Println("Number is valid")
	} else {
		fmt.Println("Number is not valid")
	}
}

func numberornot(number string) bool {
	_, err := strconv.ParseInt(number, 10, 64)
	if err == nil {
		if len(number) == 10 {
			return true
		}

	}

	return false
}
