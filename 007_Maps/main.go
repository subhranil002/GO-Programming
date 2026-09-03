package main

import "fmt"

func main() {
	// 1. Create a map
	ages := map[string]int{
		"John": 25,
		"Alex": 30,
		"Sam":  22,
	}

	fmt.Println(ages)

	// 2. Access a value using key
	fmt.Println(ages["John"])
	// 25

	// 3. Add a new key-value pair
	ages["David"] = 28

	fmt.Println(ages)

	// 4. Update a value
	ages["John"] = 26

	fmt.Println(ages["John"])
	// 26

	// 5. Delete a key
	delete(ages, "Sam")

	fmt.Println(ages)

	// CHECK IF KEY EXISTS
	age, exists := ages["Alex"]

	if exists {
		fmt.Println("Alex's age:", age)
	} else {
		fmt.Println("Alex not found")
	}

	// Key that doesn't exist
	age, exists = ages["Tom"]

	fmt.Println(age)    // 0
	fmt.Println(exists) // false

	// CREATE EMPTY MAP
	people := make(map[string]int)

	people["John"] = 25
	people["Alex"] = 30

	fmt.Println(people)

	// LENGTH
	fmt.Println("Number of entries:", len(people))

	// MAP OF SLICES
	marks := map[string][]int{
		"John": {90, 85, 95},
		"Alex": {80, 88, 92},
	}

	fmt.Println(marks["John"])
	// [90 85 95]
}
