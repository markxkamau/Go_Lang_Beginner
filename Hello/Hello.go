package main

import "fmt"

var (
	name string = "Mark"
	age int
	year int
)

func main() {
	var pntr *string

	greeting := "Greetings from "+name

	pntr = &greeting

	fmt.Println(*pntr)
	fmt.Printf("%v was born in the year %v and is currently %v years old",name, year,age)
}
