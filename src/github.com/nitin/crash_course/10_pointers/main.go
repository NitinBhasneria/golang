package main

import "fmt"

func main() {
	a := 5
	b := &a // Pointer

	fmt.Println(a, b)
	fmt.Printf("%T\n", b) // For data type of pointer

	// Use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val with pointer
	*b = 6
	fmt.Println(a)
}
