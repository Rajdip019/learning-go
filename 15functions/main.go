package main

import "fmt"

func main() {
	fmt.Println("Functions in golang")
	greeter()
	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult, myMessage := proAdder(3, 5, 10, 34, 45, 2734, 21, 4)
	fmt.Println("The pro result is: ", proResult)
	fmt.Println("Pro message is: ", myMessage)
}

func greeter() {
	fmt.Println("Hello my friend")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi Pro Result Function"
}
