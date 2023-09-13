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
	alex.print("Alex:")
	sandy.print("Sandy:")

	// Create person with default values
	var priti person
	priti.print("Priti normal:")
	priti.printf("Priti formatted:")
	priti.firstName = "Priti"
	priti.lastName = "Sharma"
	priti.contact = contactInfo{email: "priti@sharma.com", pincode: 23099}
	priti.printf("Priti with contactInfo, formatted:")

	//Try to update first name of sandy
	sandy.updateFirstName(("Cindy"))
	sandy.print("Update Sandy'S first name") //First name did not update - no error too

	// Update Name with pointers
	sandyPointer := &sandy //Memory of
	sandyPointer.updateName("Cindy", "Klara")
	sandy.print("Update Sandy's name with pointer:")

	//Go shortcut to update person without using pointer.
	// Hint: use the type pointer to a struct as the receiver
	sandy.updateName("Sandeep", "Kaur")
	sandy.print("Update Sandy's name without passing pointer")

	//Use go shortcut to update person without pointer, in func def and func body
	arun := person{"Arun", "Jaitley", contactInfo{"a@jaitley.com", 67556}}
	arun.print("Arun(init):")
	arun.updateEmail("j@arun.com")
	arun.print("Arun, udpated email without using *:")
}

// Update the email of the person through a pointer without using *
func (pointerToPerson *person) updateEmail(email string) {
	pointerToPerson.contact.email = email
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
func (p person) print(version string) {
	fmt.Println(version, p)
}

// Prints the person struct in a given format
func (p person) printf(version string) {
	fmt.Printf("%s%+v \n", version, p)
}
