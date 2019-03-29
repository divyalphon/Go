package main

import (
	"encoding/json"
	"fmt"
)

type PersonDetails struct {
	Firstname     string `json:"firstname"`
	Lastname      string `json:"lastname"`
	Profession    string `json:"profession"`
	Education     string `json:"education,omitempty"`
	NickName      string `json:"nickName"`
	Age           string `json:"Age"`
	Sibling       string `json:"Sibling"`
	FavouriteFood string `json:"Favourite Food"`
	LoveTo        string `json:"LoveTo"`
	PlaceOfBirth  string `json:"PlaceOfBirth"`
}

var persondetailsJson = []byte(`[{
	"firstname": "Irene",
	"lastname": "Varun",
	"profession": "Student",
	"education": "Ist Grade",
	"nickName": "Iru",
	"Age": "6",
	"Sibling": "1",
	"Favourite Food": "Everything Healthy",
	"LoveTo": "Sing",
	"PlaceOfBirth": "Dubai"
	},{
	"firstname": "Ariana",
	"lastname": "Varun",
	"profession": "Baby",
	"nickName": "Hola",
	"Age": "2",
	"Sibling": "1",
	"Favourite Food": "Everything Healthy",
	"LoveTo": "Climbing",
	"PlaceOfBirth": "Bangalore"
	}]
	`)

func main() {
	var persondetails []PersonDetails

	json.Unmarshal(persondetailsJson, &persondetails)

	fmt.Println(persondetails)

}
