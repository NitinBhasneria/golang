package main

import "fmt"

type I interface {
	M()
}

type person struct {
	name string
}

func main() {
	// var bob = person1{"Nitin"}
	var jane = person{"Naren"}
	// describe(bob)
	var j I
	j = jane
	describe(j)
	// describe(5)
	// describe("niitn")
}

func describe(i I) {
	i.M()
}

func (p person) M() {
	fmt.Printf("%v", p)
}
