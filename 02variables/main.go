package main

import "fmt"

const LoginToken string = "jshjksld" //Public

func main() {
	var username string = "Rajdeep"
	fmt.Println(username)
	fmt.Printf("Varibale is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Varibale is of type: %T \n", isLoggedIn)

	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("Varibale is of type: %T \n", smallVal)

	var smallFloat float64 = 255.38479825698235629835
	fmt.Println(smallFloat)
	fmt.Printf("Varibale is of type: %T \n", smallFloat)

	// default values and some aliases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("Varibale is of type: %T \n", anothervariable)

	//Implicit type

	var website = "Rajdeep is a good boy!"
	fmt.Println(website)
	fmt.Printf("Varibale is of type: %T \n", anothervariable)

	// no var style

	numberofUser := 30000
	fmt.Println(numberofUser)

	fmt.Println(LoginToken)
	fmt.Printf("Varibale is of type: %T \n", LoginToken)
}
