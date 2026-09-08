package main

import "fmt"

// 1. Basic function
func greet() {
	fmt.Println("Hello")
}

// 2. Function with parameter
func greetUser(name string) {
	fmt.Println("Hello", name)
}

// 3. Function with return value
func add(a int, b int) int {
	return a + b
}

// 4. Multiple parameters of same type
func multiply(a, b int) int {
	return a * b
}

// 5. Multiple return values
func divide(a, b int) (int, int) {
	return a / b, a % b
}

// 6. Named return values
func calculate(a, b int) (sum int, difference int) {
	sum = a + b
	difference = a - b
	return
}

// 7. Variadic function
// Accepts any number of arguments
func total(numbers ...int) int {
	sum := 0

	for _, n := range numbers {
		sum += n
	}

	return sum
}

func main() {

	// Call basic function
	greet()

	// Function with parameter
	greetUser("John")

	// Function with return value
	result := add(10, 20)
	fmt.Println(result)
	// 30

	// Multiple parameters
	fmt.Println(multiply(5, 4))
	// 20

	// Multiple return values
	quotient, remainder := divide(10, 3)

	fmt.Println(quotient)
	// 3

	fmt.Println(remainder)
	// 1

	// Ignore a return value using _
	quotient, _ = divide(10, 3)
	fmt.Println(quotient)

	// Named return values
	sum, difference := calculate(10, 5)

	fmt.Println(sum)
	// 15

	fmt.Println(difference)
	// 5

	// Variadic function
	fmt.Println(total(10, 20, 30))
	// 60

	// Pass a slice to variadic function
	numbers := []int{10, 20, 30, 40}

	fmt.Println(total(numbers...))
	// 100
}