package main

import (
	"database/sql"
	"fmt"

	_ "gopkg.in/goracle.v2"
)

func main() {
	db, err := sql.Open("goracle", "username/password@10.0.1.127:1521/orclpdb1")

	fmt.Println(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		fmt.Printf("Error connecting to the database: %s\n", err)
		return
	}

	employee_id := 8
	company_name := "XXXXXXXX"
	designation := "Engineer"
	rows, err := db.Query("select count(*) from employee_details where employee_id= :employee_id and company_name= :company_name and designation= :designation", employee_id, company_name, designation)
	if err != nil {
		fmt.Println("Error fetching addition")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var sum int
		rows.Scan(&sum)
		fmt.Println("combinationCount   ", sum)
	}

}
