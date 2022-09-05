package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Learn Slice with go...")
	var fruitList = []string{"Apple", "Tomato", "Banana"}
	fmt.Printf("Type if friuit list is %T \n", fruitList)

	fruitList = append(fruitList, "Water Melon", "Mango")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 123
	highScores[1] = 954
	highScores[2] = 169
	highScores[3] = 1208
	highScores = append(highScores, 555, 666, 777, 888)

	fmt.Println(sort.IntsAreSorted(highScores))
	fmt.Println("Sorting scores...")
	sort.Ints(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	fmt.Println("The high scroes are", highScores)

	// How to remove a value form slice based on index

	var courses = []string{"react.js", "javascript", "swift", "python", "DSA"}
	fmt.Println("All the courses are: ", courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
