package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Index is %v and the value is %v \n", index, day)
	// }

	rougeValue := 1

	for rougeValue < 10 {
		if rougeValue == 2 {
			goto rajdeep
		}
		fmt.Println("Value is", rougeValue)
		rougeValue++
	}

rajdeep:
	fmt.Println("Juping at rajdeep.space")
}
