package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Golang")

	language := make(map[string]string)

	language["JS"] = "JavaScript"
	language["RB"] = "Ruby"
	language["PY"] = "Python"
	language["TS"] = "TypeScript"

	fmt.Println("List of all languages: ", language)
	fmt.Println("Js shorts for: ", language["JS"])

	delete(language, "RB")
	fmt.Println("List of all languages: ", language)

	// Looping through the map languages
	for key, value := range language {
		fmt.Printf("For key %v value is %v\n", key, value)
	}
}
