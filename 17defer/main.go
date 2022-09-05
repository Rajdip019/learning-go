package main

import "fmt"

func main() {
	defer fmt.Println("Hello world")
	defer fmt.Println("one")
	defer fmt.Println("two")

	fmt.Println("Defer in golang")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
