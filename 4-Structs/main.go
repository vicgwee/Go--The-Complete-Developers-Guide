package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
	fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipcode: 123456,
		},
	}

	/* Proper way to do things: create a pointer and call updateName on it
	jimptr := &jim
	jimptr.updateName("Jimmy")
	*/

	//Shortcut: Go automatically converts jim into a pointer if there is a receiver with type *person
	jim.updateName("Jimmy")

	jim.print()
}

func (ptr *person) updateName(newFirstName string) {
	(*ptr).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
