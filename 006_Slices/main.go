package main

import "fmt"

func main() {
	// 1. Create a slice
	numbers := []int{10, 20, 30, 40, 50}

	fmt.Println(numbers)
	// [10 20 30 40 50]

	// 2. Access elements
	fmt.Println(numbers[0]) // 10
	fmt.Println(numbers[2]) // 30

	// 3. Change an element
	numbers[0] = 100

	fmt.Println(numbers)
	// [100 20 30 40 50]

	// 4. Length of slice
	fmt.Println("Length:", len(numbers))
	// 5

	// 5. Append elements
	numbers = append(numbers, 60)

	fmt.Println(numbers)
	// [100 20 30 40 50 60]

	// Append multiple elements
	numbers = append(numbers, 70, 80, 90)

	fmt.Println(numbers)
	// [100 20 30 40 50 60 70 80 90]


	// CREATE EMPTY SLICE
	var names []string

	fmt.Println(names)
	// []

	fmt.Println(len(names))
	// 0

	names = append(names, "John")
	names = append(names, "Alex")

	fmt.Println(names)
	// [John Alex]


	// MAKE()
	// make(type, length)
	scores := make([]int, 5)

	fmt.Println(scores)
	// [0 0 0 0 0]

	fmt.Println(len(scores))
	// 5


	// make(type, length, capacity)
	values := make([]int, 3, 5)

	fmt.Println(values)
	fmt.Println("Length:", len(values))
	fmt.Println("Capacity:", cap(values))


	// SLICE FROM AN ARRAY
	array := [5]int{10, 20, 30, 40, 50}

	s := array[1:4]

	fmt.Println(s)
	// [20 30 40]


	// SLICE SYNTAX
	fmt.Println(array[:3])
	// [10 20 30]

	fmt.Println(array[2:])
	// [30 40 50]

	fmt.Println(array[:])
	// [10 20 30 40 50]


	// COPY
	a := []int{1, 2, 3}

	b := make([]int, len(a))

	copy(b, a)

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	
	// REMOVE ELEMENT
	items := []string{"A", "B", "C", "D"}

	// Remove element at index 1 ("B")
	items = append(items[:1], items[2:]...)

	fmt.Println(items)
	// [A C D]
}
