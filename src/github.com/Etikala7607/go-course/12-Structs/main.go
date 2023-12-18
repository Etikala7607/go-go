package main

import "fmt"

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	age       int
	gender    string
}

// Greeting method value receiver
func (p Person) greeting() string {
	return "Hello" + p.firstName + " " + p.lastName + " and " + fmt.Sprint(p.age) + " from " + p.city
}

//has birthday method pointer receiver
func (p *Person) hasBirthday() {
	p.age++
}

//getsMarried pointer reciver
func (p *Person) getMarried(spouseLastName string) {

	if p.gender == "f" {
		return
	} else {
		p.lastName = spouseLastName
	}
}
func main() {
	// init person using struct
	person1 := Person{firstName: "Samantha", lastName: "ssgge", age: 23, city: "lolooo", gender: "m"}
	//Alternative
	//person2 := Person{"Ravan", "sssrs", "fsgsfs", 25, "fhfj"}
	//fmt.Println(person1)
	//fmt.Println(person2)
	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)
	person1.hasBirthday()
	person1.getMarried("sravan")
	fmt.Println(person1.greeting())

}
