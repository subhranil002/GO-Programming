package main

import "fmt"

// Private Variable (Var name starts with small case) (Can not be accessable from other packages)
var prevUser = "Cooper"

// Public Variable (Var name starts with capital case) (Can be accessable from other packages)
var CurrUser = "Subhranil"

func main() {
	var username string = "Subhranil"
	fmt.Println(username)
	fmt.Printf("Variable is type of %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is type of %T \n", isLoggedIn)

	var numberOfOrders int32 = 500
	fmt.Println(numberOfOrders)
	fmt.Printf("Variable is type of %T \n", numberOfOrders)

	var marksInLastExam float32 = 76.55
	fmt.Println(marksInLastExam)
	fmt.Printf("Variable is type of %T \n", marksInLastExam)

	// Default Values

	var defStr string
	fmt.Println(defStr)

	var defBool bool
	fmt.Println(defBool)
	
	var defInt int
	fmt.Println(defInt)
	
	var defFloat float32
	fmt.Println(defFloat)

	// Implicit type

	var initials = "SC"
	fmt.Println(initials)
	fmt.Printf("Variable is type of %T \n", initials)

	// No var style (Inside method only)

	noOfUser := 30000
	fmt.Println(noOfUser)

	// Constant Variable

	const adminName = "Admin"
	fmt.Println(adminName)
}
