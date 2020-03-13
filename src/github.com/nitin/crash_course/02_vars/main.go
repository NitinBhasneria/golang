package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	//int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	var name = "Brad"
	var age int = 37

	// Using Shorthand
	// name2 := "Brad2"
	// size := 1.3

	name2, email := "Brad2", "brad@email.com"

	fmt.Println(name, age, name2, email)
	fmt.Printf("%T\n", age)
}
