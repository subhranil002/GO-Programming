package main

import "fmt"

// BASIC STRUCT
type Person struct {
	Name string
	Age  int
	City string
}


// STRUCT WITH DIFFERENT TYPES
type Student struct {
	Name   string
	Age    int
	Marks  float64
	Passed bool
}


// NESTED STRUCT
type Address struct {
	City  string
	State string
}
type Employee struct {
	Name    string
	Age     int
	Address Address
}

// ATTACH METHOD TO STRUCT
func (p Person) greet() {
	fmt.Println("Hello", p.Name)
}

// FUNCTION ACCEPTING STRUCT
func printPerson(p Person) {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("City:", p.City)
}


func main() {

	// Create a struct with zero values
	var p Person

	fmt.Println(p)
	// { 0 }

	// Assign values
	p.Name = "John"
	p.Age = 25
	p.City = "Bangalore"

	fmt.Println(p)
	// {John 25 Bangalore}

	// Access fields
	fmt.Println(p.Name)
	// John

	fmt.Println(p.Age)
	// 25

	fmt.Println(p.City)
	// Bangalore

	// Update a field
	p.Age = 26

	fmt.Println(p.Age)
	// 26

	// Struct initialization using field names
	person := Person{
		Name: "Alex",
		Age:  30,
		City: "Mumbai",
	}

	fmt.Println(person)

	// Positional initialization
	person2 := Person{"Sam", 22, "Delhi"}

	fmt.Println(person2)

	// Struct with different data types
	student := Student{
		Name:   "John",
		Age:    20,
		Marks:  85.5,
		Passed: true,
	}

	fmt.Println(student)
	fmt.Println(student.Name)
	fmt.Println(student.Marks)
	fmt.Println(student.Passed)

	// Struct containing a slice
	type StudentMarks struct {
		Name  string
		Marks []int
	}
	sm := StudentMarks{
		Name:  "John",
		Marks: []int{90, 85, 95},
	}

	fmt.Println(sm)
	fmt.Println(sm.Marks)
	fmt.Println(sm.Marks[0])

	// Nested struct
	employee := Employee{
		Name: "David",
		Age:  28,
		Address: Address{
			City:  "Bangalore",
			State: "Karnataka",
		},
	}

	fmt.Println(employee)

	// Access nested fields
	fmt.Println(employee.Name)
	fmt.Println(employee.Address.City)
	fmt.Println(employee.Address.State)

	// Multiple struct variables
	p1 := Person{
		Name: "John",
		Age:  25,
		City: "Bangalore",
	}
	p2 := Person{
		Name: "John",
		Age:  25,
		City: "Bangalore",
	}

	fmt.Println(p1)
	fmt.Println(p2)

	// Compare structs
	fmt.Println(p1 == p2)
	// true

	// Function accepting a struct
	printPerson(p1)

	// Attached method call
	p1.greet()
	// Hello John
}