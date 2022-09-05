package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with files in go!")
	content := "This need to go in a file - rajdeep.space"

	file, err := os.Create("./mywebsitefile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mywebsitefile.txt")
}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println("text data inside the file is: \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
