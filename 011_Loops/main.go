package main

import "fmt"

func main() {

	// 1. Basic for loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// 2. while-style loop
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// 3. Infinite loop
	/*
	for {
		fmt.Println("Running")
	}
	*/

	// 4. break
	// Stops the loop completely
	for i := 1; i <= 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
	// 1 2

	// 5. continue
	// Skips current iteration
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
	// 1 2 4 5
}