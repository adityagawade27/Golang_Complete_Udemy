package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func main() {
	// alex := person{"Alex", "Anderson"} --> correct but very confusing
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	jim := person{
		firstName: "Jim",
		lastName:  "Morrison",
		contact: contactInfo{
			email:   "jim@example.com",
			zipCode: 85027,
		},
	}

	jim.updateName("jimmy")
	jim.print()

	// Print out field names and values
	//fmt.Printf("%+v", alex)

}
