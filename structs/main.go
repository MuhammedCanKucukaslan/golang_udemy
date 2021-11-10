package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo contactInfo
}

func (p person) print() {
	fmt.Print("p: ", p.firstName, ", ", p.lastName, ", ", p.contactInfo.email, ", ", p.contactInfo.zipcode, ".\n")

}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {
	p := person{"name", "a", contactInfo{"mck@gmail.com", 94000}}
	p.print()

	pPtr := &p
	pPtr.updateName("lol")
	p.print()
	(&p).updateName("evenMoreLoL")
	p.print()
	p.updateName("type tf")
	p.print()
}
