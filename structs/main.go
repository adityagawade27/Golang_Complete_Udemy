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

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
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

	//jimPointer := &jim
	// In go You can use either jim.updateName or jimPointer.updateName to refer to a function
	// that recieves the pointer as a reciever.
	// Go will automatically resolve the first case to a pointer.
	// You will need to make pointer function if you need to update values in the existing struct
	// Go will copy the struct to a new memory location and run the func there only without pointers.
	jim.updateName("jimmy")
	jim.print()
	// Important point with go -- if we define a reciever, pointer,
	// Print out field names and values
	//fmt.Printf("%+v", alex)

}
