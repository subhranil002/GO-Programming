package main

import "fmt"

func main() {
	// Type conversion: int to float64
	var a int = 10
	var b float64 = float64(a)

	fmt.Println(a)
	fmt.Println(b)

	// float64 to int
	var x float64 = 10.5
	var y int = int(x)

	fmt.Println(y) // 10 (decimal part is removed)

	// int to string using conversion
	var n int = 65
	var s string = string(rune(n))

	fmt.Println(s) // A

	// String to []byte
	str := "Hello"
	bytes := []byte(str)

	fmt.Println(bytes)

	// []byte to string
	newStr := string(bytes)

	fmt.Println(newStr)
}