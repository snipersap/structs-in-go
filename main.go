package main

import "fmt"

//Declare struct
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// Declare sub struct
type contactInfo struct {
	email   string
	pincode int
}

func main() {
	// Create person values
	alex := person{"Alex", "Hardy", contactInfo{"h@h.com", 99888}} // Order of the field is important, all fields must
	sandy := person{firstName: "Sandy", lastName: "Homie", contact: contactInfo{"sandy.homie@gmail.com", 788999}}

	// Print person values to the screen
	alex.print()
	sandy.print()

	// Create person with default values
	var priti person
	priti.print()
	priti.printf()
	priti.firstName = "Priti"
	priti.lastName = "Sharma"
	priti.contact = contactInfo{email: "priti@sharma.com", pincode: 23099}
	priti.printf()

}

func (p person) print() {
	fmt.Println(p)
}

func (p person) printf() {
	fmt.Printf("%+v \n", p)
}
