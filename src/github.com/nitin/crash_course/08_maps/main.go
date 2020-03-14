package main

import (
	"fmt"
)

func main() {
	// Define map
	emails := make(map[string]string)

	// Assign
	emails["Bob"] = "bob@gmail.com"
	emails["Brad"] = "Brad@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Brad"])

	// Delete
	delete(emails, "Bob")
	fmt.Println(emails)

	//Declare amp and add kv
	names := map[string]string{"Nitin": "Bhasneria", "Muskan": "Choudhary"}

	fmt.Println(names)
}
