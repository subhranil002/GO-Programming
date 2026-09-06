package main

import "fmt"

func main() {

    day := 2

    // Basic switch
    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    default:
        fmt.Println("Invalid day")
    }

    // Multiple values in one case
    switch day {
    case 1, 7:
        fmt.Println("Weekend")
    default:
        fmt.Println("Weekday")
    }

    // Switch without an expression
    age := 20

    switch {
    case age < 13:
        fmt.Println("Child")
    case age < 18:
        fmt.Println("Teenager")
    default:
        fmt.Println("Adult")
    }
}