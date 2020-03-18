package main

import "fmt"

type I interface {
	M()
}

type person struct {
	name string
}

func main() {
	var bob = person{"Nitin"}
	var jane = person{"Naren"}
	describe(bob)
	describe(jane)
	describe(5)
	describe("Niitn")
	// describe(5)
	// describe("niitn")
}

func describe(i interface{}) {
	fmt.Printf("%v\n", i)
}
