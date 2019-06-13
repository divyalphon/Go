package main

import "fmt"

type Persons struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	people := []Persons{{"Divya", 12}, {"Divya", 12}, {"Irene", 10}}
	fmt.Println(RemoveDuplicatesFromSaolaErrorDetails(people))
}

func RemoveDuplicatesFromSaolaErrorDetails(s []Persons) []Persons {
	m := make(map[Persons]bool)
	for _, item := range s {
		if _, ok := m[item]; ok {
			// duplicate item

		} else {
			m[item] = true
		}
	}

	var result []Persons
	for item, _ := range m {
		result = append(result, item)
	}
	return result
}
