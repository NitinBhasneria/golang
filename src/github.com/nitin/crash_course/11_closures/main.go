package main

import "fmt"

// Anonymous function
func returnAnony() func(string) {
	return func(x string) {
		fmt.Println(x)
	}
}

// Anonymouse function forming closures
func addr() func(int) int {
	sum := 0
	// Here it forms a closure as it remembers
	// and has a access to the local variables
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	printAnony := returnAnony()
	printAnony("Hello it worked")
	sum := addr()
	for i := 0; i <= 10; i++ {
		fmt.Println(sum(i))
	}
}
