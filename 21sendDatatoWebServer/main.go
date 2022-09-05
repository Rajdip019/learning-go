package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Sending request to web server.")
	PerformGetRequest()
	performPostRequest()
	PerformFormRequest()
}

// Performing a get request
func PerformGetRequest() {
	const myurl = "http://localhost:1111/get"

	res, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Status code", res.StatusCode)
	fmt.Println("Content length is: ", res.ContentLength)

	var resString strings.Builder
	content, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(content))

	byteCount, _ := resString.Write(content)
	fmt.Println("Byte Count is", byteCount)
	fmt.Println(resString.String())

}

// Performing a post request
func performPostRequest() {
	const myurl = "http://localhost:1111/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"name" : "Rajdeep Sengupta",
			"tech" : "Golang",
			"age" : 20,
			"social" : "linkedIn"
		}
	`)

	res, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var resString strings.Builder
	content, _ := ioutil.ReadAll(res.Body)
	byteCount, _ := resString.Write(content)
	fmt.Println("The byte count is: ", byteCount)
	fmt.Println(resString.String())
}

func PerformFormRequest() {
	const myurl = "http://localhost:1111/postform"

	//formdata

	data := url.Values{}
	data.Add("firstName", "Rajdeep")
	data.Add("lastName", "Sengupta")
	data.Add("email", "rajdeep@go.dev")
	// data.Add("age", 20)

	res, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var resString strings.Builder
	content, _ := ioutil.ReadAll(res.Body)

	byteCount, _ := resString.Write(content)
	fmt.Println("The byte count is: ", byteCount)
	fmt.Println(resString.String())
}
