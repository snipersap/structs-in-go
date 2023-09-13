package main

import (
	"testing"
)

func TestUpdateName(t *testing.T) {
	// Setup
	johnDoe := person{"John", "Doe", contactInfo{"j@doe.com", 11111}}
	expectedEmail := "doe@john.com"
	// Execute
	johnDoe.updateEmail(expectedEmail)
	if johnDoe.contact.email != expectedEmail {
		t.Errorf("Expected the email of John to be:%s, found:%s instead.", expectedEmail, johnDoe.contact.email)
	}
	//Teardown
	johnDoe = person{}
	expectedEmail = ""
}

func TestUpdateFirstName(t *testing.T) {
	// Setup
	maryAnne := person{firstName: "Mary", lastName: "Anne", contact: contactInfo{email: "m@anne.com", pincode: 22222}}
	expectedFirstName := maryAnne.firstName
	// Execute
	maryAnne.updateFirstName("Meryll")
	if maryAnne.firstName != expectedFirstName {
		t.Errorf("Expected first name to be %s, found %s", expectedFirstName, maryAnne.firstName)
	}
}
