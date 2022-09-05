package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers!")

	// var ptr *int
	// fmt.Println("This is a pointer", ptr)

	myNumber := 19

	var ptr = &myNumber
	fmt.Println("The pointer value is", ptr)
	fmt.Println("The pointer value is", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value is : ", myNumber)
}
