package main

import "fmt"

//Declare struct
type person struct {
	firstName string
	lastName  string
}

func main() {
	// Create person values
	alex := person{"Alex", "Hardy"} // Order of the field is important
	sandy := person{firstName: "Sandy", lastName: "Homie"}

	// Print person values to the screen
	alex.print()  //{Alex Hardy}
	sandy.print() //{Sandy Homie}

	// Create person with default values
	var priti person
	priti.print()  //{ }
	priti.printf() //{firstName: lastName:}
	priti.firstName = "Priti"
	priti.lastName = "Sharma"
	priti.printf() //{firstName:Priti lastName:Sharma}

}

func (p person) print() {
	fmt.Println(p)
}

func (p person) printf() {
	fmt.Printf("%+v \n", p)
}
