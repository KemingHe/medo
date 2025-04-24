package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// brutus := person{"Brutus", "Buckeye"} // Avoid using implicitly-ordered fields to declare new struct
	// brutus := person{firstName: "Brutus", lastName: "Buckeye"} // This is perferred

	// var brutus person
	// brutus.firstName = "Brutus"
	// brutus.lastName = "Buckeye"

	carmen := person{
		firstName: "Carmen",
		lastName:  "Canvas",
		contactInfo: contactInfo{
			email:   "canvas.1@osu.edu",
			zipCode: 43210,
		}, // Go requires trailing comma for multi-line composite literals, i.e struct
	}

	// carmenPtr := &carmen // &var returns the memory address of var, & is an operator
	// carmenPtr.updateFirstName("Carlos")

	carmen.updateFirstName("Carlos")
	carmen.print() 

	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

// Go pass by value, so this is incorrect and will create a separate person struct to modify the firstName property
// the resulting change will not be reflected on the original struct
// func (p person) updateFirstName(newFirstName string) {
// 	p.firstName = newFirstName
// }

// Go shortcut: pointer receiver accepts both pointer and original variable
func (ptr *person) updateFirstName(newFirstName string) { // *type represents pointer to a var of type, * here is part of type description
	(*ptr).firstName = newFirstName // *ptr returns the value stored at memory address prt, * is an operator
}

func (p person) print() {
	fmt.Printf("%+v\n", p) // %+v prints out all field names and values of struct
}

// The * operator can turn ptr (address) into its corresponding value, i.e. *ptr
// The & operator can turn var (value) into its corresponding address in memory, i.e. &var

// This works without ptr because slice is already a reference type in Go
func updateSlice(s []string) {
	s[0] = "Good morning"
}

// Go value types (need to use ptr for mutation): int, float, string, bool, struct, etc.
// Go reference types: slice, map, channel, pointer, function, etc.
