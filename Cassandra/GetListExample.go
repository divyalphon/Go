package main

import (
	"fmt"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

//for get example ,you dont have to worry about cql json ,it is required for only inserting into cassandra
//to create type in cassandra
// CREATE TYPE keyspace.ParentsName (
//     FathersName text,
//     MothersName text,
//     Siblings int,
//     );

//while creating table PersonDetails
//to create  familydetails frozen<ParentsName>,
type PersonDetails struct {
	Firstname     string      `json:"firstname"`
	Lastname      string      `json:"lastname"`
	Profession    string      `json:"profession"`
	Education     string      `json:"education,omitempty"`
	NickName      string      `json:"nickName"`
	FamilyDetails ParentsName `json:"familyDetails"`
	Age           string      `json:"Age"`
	Sibling       string      `json:"Sibling"`
	FavouriteFood string      `json:"Favourite Food"`
	LoveTo        string      `json:"LoveTo"`
	PlaceOfBirth  string      `json:"PlaceOfBirth"`
}

type ParentsName struct {
	FathersName int    `json:"fathersName"`
	MothersName string `json:"mothersName"`
	Siblings    int    `json:"siblings"`
}

func init() {
	var err error

	cluster := gocql.NewCluster("127.0.0.1")
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}

}

func main() {
	CassandraSession := Session
	defer CassandraSession.Close()
	m := map[string]interface{}{}
	rowValues := make(map[string]interface{})
	var d PersonDetails
	var dd []PersonDetails

	iter := Session.Query("SELECT * FROM keyspace.PersonDetails where person_id='12345' allow filtering").Iter()
	for iter.MapScan(rowValues) {
		d = PersonDetails{
			Firstname:     m["firstname"].(string),
			Lastname:      m["lastname"].(string),
			Profession:    m["profession"].(string),
			Education:     m["education"].(string),
			NickName:      m["nickName"].(string),
			FamilyDetails: m["familyDetails"].(ParentsName),
			Age:           m["age"].(string),
			Sibling:       m["sibling"].(string),
			FavouriteFood: m["favouriteFood"].(string),
			LoveTo:        m["loveTo"].(string),
			PlaceOfBirth:  m["placeOfBirth"].(string),
		}
		dd = append(dd, d)
		rowValues = make(map[string]interface{})
	}

	fmt.Println("rows  ", dd)

}
