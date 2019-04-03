package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//divya is the name of waitgroup
var divya sync.WaitGroup

func main() {
	rand.Seed(time.Now().Unix())
	projects := make(chan string, 0)
	divya.Add(5) //adding 5 goroutines to handle the channel projects
	for i := 1; i <= 5; i++ {
		go employee(projects, i)
	}
	for j := 1; j <= 10; j++ {
		projects <- fmt.Sprintf("Project :%d", j)
	}

	// Close the channel so the goroutines will quit
	close(projects)
	divya.Wait()
}

func employee(projects chan string, i int) {
	defer divya.Done()
	for {
		project, result := <-projects
		if result == false {
			//this means the channel is empty and closed
			fmt.Println("*******************  ", employee)
			return
		}
		fmt.Printf("Employee : %d : Started   %s\n", employee, project)

		// Randomly wait to simulate work time.
		sleep := rand.Int63n(50)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		// Display time to wait
		fmt.Println("\nTime to sleep", sleep, "ms\n")

		// Display project completed by employee.
		fmt.Printf("Employee : %d : Completed %s\n", employee, project)
	}
}
