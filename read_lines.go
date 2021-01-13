package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./file.txt")

	if err != nil {
		fmt.Printf("There was an error trying to open the file")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
	}
}
