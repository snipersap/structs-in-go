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

// Entry to structs-in-go
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

	//Try to update first name of sandy
	sandy.updateFirstName(("Cindy"))
	sandy.print() //First name did not update - no error too

	// Update Name with pointers
	sandyPointer := &sandy //Memory of
	sandyPointer.updateName("Cindy", "Klara")
	sandy.print()

	//Go shortcut to update person without using pointer.
	// Hint: use the type pointer to a struct as the receiver
	sandy.updateName("Sandeep", "Kaur")
	sandy.print()
}

// Update the name of the person through a pointer
func (pointerToPerson *person) updateName(newFirstName string, newLastName string) {
	(*pointerToPerson).firstName = newFirstName //Value at (*) pointerToPerson
	(*pointerToPerson).lastName = newLastName
}

// Tries to update the first name of the struct person
func (p person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}

// Prints the person struct directly
func (p person) print() {
	fmt.Println(p)
}

// Prints the person struct in a given format
func (p person) printf() {
	fmt.Printf("%+v \n", p)
}
