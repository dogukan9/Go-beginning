package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {

	// alex:=person{"Alex","Anderson"}

	// alex:=person{firstName:"Alex",lastName: "Anderson"}

	// var alex person
	// alex.firstName="Alex"
	// alex.lastName="Anderson"

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jimPointer := &jim //take the address of jim
	jimPointer.updateName("jimmy")
	jim.print()

	jim.updateName("jimmy") // still is gonna work because updateName takes *person means pointer so this is a shortcut

	jim.print()
}

func (p person) print() {

	fmt.Printf("%+v", p)

}

func (pointerToPerson *person) updateName(newFirstName string) {
	// *person is just for declaring the type (type description)

	//*pointerToPerson means we are changing the spesific adress value
	(*pointerToPerson).firstName = newFirstName

}
