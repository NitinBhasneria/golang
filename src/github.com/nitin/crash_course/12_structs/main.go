package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// fistrName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	fistrName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, " + p.fistrName + " age : " + strconv.Itoa(p.age)
}

// hasBirth method (pointer reciever)
func (p *Person) hasBirth() {
	p.age++
}

func main() {
	// Init person using struct
	person1 := Person{fistrName: "Nitin", lastName: "Bhasneria", city: "Sjp",
		gender: "M", age: 20}
	// Alternate
	//person1 := Person{"Nitin", "Bhasneria", "Sjp", "M", 20}
	fmt.Println(person1)
	fmt.Println(person1.fistrName)
	person1.age++
	fmt.Println(person1.age)
	person1.hasBirth()
	fmt.Println(person1.greet())

}
