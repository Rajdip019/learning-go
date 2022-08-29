package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of Golang!")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006"))
}
