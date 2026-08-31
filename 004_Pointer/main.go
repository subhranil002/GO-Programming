package main

import "fmt"

func main() {
	// 1. Normal variable
	x := 10

	fmt.Println("Value of x:", x)
	// 10


	// 2. & → Address operator
	// &x gives the memory address of x

	fmt.Println("Address of x:", &x)


	// 3. * → Pointer declaration
	// p stores the address of x

	var p *int = &x

	fmt.Println("Pointer p:", p)
	// Address of x


	// 4. * → Dereference operator
	// *p gives the value stored at the address

	fmt.Println("Value using pointer:", *p)
	// 10


	// 5. Changing value using pointer

	*p = 20

	fmt.Println("New value of x:", x)
	// 20

	// NIL POINTER
	var ptr2 *int

	fmt.Println("Nil pointer:", ptr2)
	// <nil>
}
