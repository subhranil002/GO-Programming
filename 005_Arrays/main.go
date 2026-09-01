package main

import "fmt"

func main() {
	// 1. Declare an array
	var numbers [5]int

	fmt.Println(numbers)
	// [0 0 0 0 0]


	// 2. Assign values using index
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println(numbers)
	// [10 20 30 40 50]


	// 3. Access an element
	fmt.Println(numbers[0]) // 10
	fmt.Println(numbers[2]) // 30


	// 4. Change an element
	numbers[0] = 100

	fmt.Println(numbers)
	// [100 20 30 40 50]


	// 5. Array initialization
	var marks = [5]int{90, 80, 70, 85, 95}

	fmt.Println(marks)


	// 6. Short declaration
	names := [3]string{"John", "Alex", "Sam"}

	fmt.Println(names)


	// 7. Let Go calculate the size
	// [...] means compiler determines the length

	cities := [...]string{"Bangalore", "Mumbai", "Delhi"}

	fmt.Println(cities)


	// 8. Find array length
	fmt.Println(len(numbers))
	// 5


	// MULTI-DIMENSIONAL ARRAY
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(matrix)

	fmt.Println(matrix[0][1])
	// 2
}