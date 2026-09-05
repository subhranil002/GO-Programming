package main

import "fmt"

func main() {

    age := 20

    // if
    if age >= 18 {
        fmt.Println("Adult")
    }

    // if else
    if age >= 18 {
        fmt.Println("Adult")
    } else {
        fmt.Println("Minor")
    }

    // if else if
    if age < 13 {
        fmt.Println("Child")
    } else if age < 18 {
        fmt.Println("Teenager")
    } else {
        fmt.Println("Adult")
    }

    // if with condition variable
    if n := 10; n > 5 {
        fmt.Println("Greater")
    }
}