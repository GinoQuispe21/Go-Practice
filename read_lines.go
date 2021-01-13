package main

import (
	"bufio"
	"fmt"
	"os"
)

func read_lines_file() bool {
	file, err := os.Open("./filex.txt")

	defer func() {
		file.Close()
		fmt.Println("Defer")

		r := recover()
		fmt.Println(r)

	}() //-> this function executes after the function returns something

	if err != nil {
		panic(err)
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

func execute_read_file() {
	execution := read_lines_file()
	fmt.Println(execution)
}

func main() {
	execute_read_file()
}
