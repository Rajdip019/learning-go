package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of Golang!")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.July, 19, 23, 23, 0, 0, time.UTC)
	fmt.Println("The Date is", createdDate)
	fmt.Println("The formatted date is: ", createdDate.Format("02-01-2006 15:04:05 Monday"))
}
