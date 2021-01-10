//Concurrent programming

package main

import (
	"fmt"
	"strings" //contain functions like split
	"time"
)

func main() {
	go my_name_slow("Gino")
	fmt.Println("I'm wait")

	go func() {
		var wait string
		fmt.Scanln(&wait)
	}()
}

func my_name_slow(name string) {
	letters := strings.Split(name, "")

	for _, letter := range letters {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letter)
	}
}
