package main

import (
	"bufio"
	"fmt"
	"os"
)

func read_lines_file() bool {
	file, err := os.Open("./file.txt")

	defer func() {
		file.Close()
		fmt.Println("Defer")
	}() //-> this function executes after the function returns something

	if err != nil {
		fmt.Printf("There was an error trying to open the file")
	}

	scanner := bufio.NewScanner(file)

	var i int

	for scanner.Scan() {
		i++
		line := scanner.Text()
		fmt.Println(i)
		fmt.Println(line)
	}

	if true {
		return true
	}

	//file.Close()
	return true
}

func main() {
	execution := read_lines_file()
	fmt.Println(execution)
}
