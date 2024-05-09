package main

import "fmt"

//* Structs can be embedded

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToDunphy *person) updateName(newName string) {
	(*pointerToDunphy).firstName = newName
}
func main() {

	alex := person{firstName: "Alex", lastName: "Dunphy", contact: contactInfo{email: "mmt@mmt.com", zipCode: 467890}}

	dunphyPointer := &alex

	dunphyPointer.updateName("Haley")
	alex.print()

}
