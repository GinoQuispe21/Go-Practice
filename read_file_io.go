package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file_data, err := ioutil.ReadFile("./file.txt")

	if err != nil {
		fmt.Printf("There was an error trying to open the file")
	}

	fmt.Println(string(file_data))
}
