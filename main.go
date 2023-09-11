package main

import "fmt"

//Declare struct
type person struct {
	firstName string
	lastName  string
}

func main() {
	// Create person values
	alex := person{"Alex", "Hardy"}
	sandy := person{firstName: "Sandy", lastName: "Homie"}

	// Print person values to the screen
	alex.print()
	sandy.print()
}

func (p person) print() {
	fmt.Println(p)
}
