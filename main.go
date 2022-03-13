package main

import (
	"fmt"

	"example.com/assignment/netflix"
	"example.com/assignment/ookla"
)

func main() {

	ooklaResult := ookla.RunMe()
	fmt.Println("Results for ookla")
	fmt.Println(ooklaResult)

	fmt.Println("Result for netflix")
	netflixResult := netflix.RunMe()
	fmt.Println(netflixResult)
}
