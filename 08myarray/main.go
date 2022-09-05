//Array is the most less utilized Data Structure in golang!

package main

import "fmt"

func main() {
	fmt.Println("Welcome to go arrays!")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[0] = "Apple"
	fruitList[3] = "Orange"

	fmt.Println("Fruit List is", fruitList)
	fmt.Println("Fruit List is", len(fruitList))

}
