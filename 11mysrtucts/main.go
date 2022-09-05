package main

import (
	"fmt"
)

func main() {
	fmt.Println("Stucts in golang")
	//There is no inheritance of golang; No super or parent

	rajdeep := User{"Rajdeep", "rajdeep@go.dev", true, 16}
	fmt.Println(rajdeep)
	fmt.Printf("Rajdeep Details are: %+v \n", rajdeep)
	fmt.Printf("Name is: %v and email is %v.", rajdeep.Name, rajdeep.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
