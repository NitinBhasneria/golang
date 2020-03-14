package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	//Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and assign
	friends := [2]string{"Nitin", "Gupta"}

	//	Slices
	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])
	fmt.Println(friends)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])
}
