package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string   `json:"course_name"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Bit more into json with GO!")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []courses{
		{"reactJs Bootcamp", 299, "learnCodeonline.in", "abc123", []string{"web-dev", "js"}},
		{"Mern Bootcamp", 199, "learnCodeonline.in", "bcd123", []string{"mern", "js"}},
		{"Angular Bootcamp", 599, "learnCodeonline.in", "hit123", []string{"angular", "js", "google"}},
		{"Js Bootcamp", 699, "learnCodeonline.in", "ok123", nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
			"course_name": "Angular Bootcamp",
			"price": 599,
			"website": "learnCodeonline.in",
			"tags": [
					"angular",
					"js",
					"google"
			]
		}
	`)

	var lcoCourse courses

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v \n", lcoCourse)
	} else {
		fmt.Println("JSOn is valid")
	}

	// Some cases where you want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v \n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is %T \n", k, v, v)
	}
}
