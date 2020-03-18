package main

import "fmt"

type I interface {
	M()
}

type person struct {
	name string
}

func main() {
	var i I
	// var bob = person{"Nitin"}
	describe(i)
	// describe(bob)
	// describe(5)
	// describe("niitn")
}

func describe(i I) {
	fmt.Printf("%v", i)
}
