package main

import (
	"fmt"
)

func main() {
	fmt.Println("Methods in golang")
	//There is no inheritance of golang; No super or parent

	rajdeep := User{"Rajdeep", "rajdeep@go.dev", true, 16}
	fmt.Println(rajdeep)
	fmt.Printf("Rajdeep Details are: %+v \n", rajdeep)

	rajdeep.GetStatus()
	rajdeep.NewMail()
	fmt.Printf("Name is: %v and email is %v.\n", rajdeep.Name, rajdeep.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active : ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is,", u.Email)
}
