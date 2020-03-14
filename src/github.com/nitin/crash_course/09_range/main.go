package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 5, 6, 5, 2}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Range with map
	names := map[string]string{"Nitin": "Bhasneria", "Muskan": "Choudhary"}

	for k, v := range names {
		fmt.Printf("%s: %s\n", k, v)
	}
}
