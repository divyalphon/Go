package main

import "fmt"

type Person struct {
	Name       string  `json:"name"`
	Profession *string `json:",omitempty"`
	Age        int     `json:"age"`
}

func main() {
	person := Person{}
	person.Age = 15
	person.Name = "Divya"
	if person.Profession == nil {
		fmt.Println(" Divya is Awesome ")
	} else {
		fmt.Println("oh shit")
	}
}
